#!/usr/bin/env bash

GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

PASS_COUNT=0
FAIL_COUNT=0

pass() {
    echo -e "${GREEN}[PASS]${NC} $1"
    ((PASS_COUNT++))
}

fail() {
    echo -e "${RED}[FAIL]${NC} $1"
    ((FAIL_COUNT++))
}

echo -e "${YELLOW}--- 🎨 Μορφοποίηση και απλοποίηση κώδικα (Standardization) ---${NC}"
if ! command -v goimports &> /dev/null; then
    fail "goimports δεν βρέθηκε! Τρέξε: go install golang.org/x/tools/cmd/goimports@latest"
else
    if goimports -w .; then
        pass "goimports -w ."
    else
        fail "goimports -w ."
    fi
fi

if gofmt -s -w .; then
    pass "gofmt -s -w ."
else
    fail "gofmt -s -w ."
fi

if ! command -v gofumpt &> /dev/null; then
    fail "gofumpt δεν βρέθηκε! Τρέξε: go install mvdan.cc/gofumpt@latest"
else
    if gofumpt -l -w .; then
        pass "gofumpt -l -w ."
    else
        fail "gofumpt -l -w ."
    fi
fi

echo -e "\n${YELLOW}--- 📦 Εκκαθάριση και επαλήθευση modules (YAGNI & KISS) ---${NC}"
if go mod tidy; then pass "go mod tidy"; else fail "go mod tidy"; fi
if go mod verify; then pass "go mod verify"; else fail "go mod verify"; fi

echo -e "\n${YELLOW}--- 📚 Έλεγχος Standard Library (YAGNI) ---${NC}"
CURRENT_MODULE=$(go list -m 2>/dev/null)
if [ -n "$CURRENT_MODULE" ]; then
    # Λίστα με επιτρεπόμενα εξωτερικά πακέτα (Whitelist).
    # Πρόσθεσε εδώ πακέτα που είναι απολύτως απαραίτητα (π.χ. "github.com/mattn/go-sqlite3").
    ALLOWED_DEPS=(
        # "github.com/google/uuid"
    )

    # Εντοπισμός πακέτων που δεν είναι Standard Library και δεν ανήκουν στο τρέχον module
    NON_STD_DEPS=$(go list -f '{{if not .Standard}}{{.ImportPath}}{{end}}' -deps ./... 2>/dev/null | grep -v "^$" | grep -v "^${CURRENT_MODULE}")
    
    # Αφαίρεση των επιτρεπόμενων πακέτων από τη λίστα των ευρημάτων
    for dep in "${ALLOWED_DEPS[@]}"; do
        NON_STD_DEPS=$(echo "$NON_STD_DEPS" | grep -v "^${dep}")
    done

    if [ -z "$NON_STD_DEPS" ]; then
        pass "Χρήση αποκλειστικά της Standard Library (ή επιτρεπόμενων εξαιρέσεων)"
    else
        fail "Εντοπίστηκαν εξωτερικά πακέτα (Επιτρέπεται μόνο η Standard Library!)"
        echo -e "  👉 ${YELLOW}Οδηγίες Διαχείρισης:${NC}"
        echo -e "  1. Βρέθηκαν τα εξής μη-standard πακέτα:\n${RED}${NON_STD_DEPS}${NC}"
        echo -e "  2. Αν αυτά τα πακέτα είναι απολύτως απαραίτητα, πρόσθεσέ τα στον πίνακα ${YELLOW}ALLOWED_DEPS${NC} του check.sh."
        echo -e "  3. Αλλιώς, συμβουλέψου το ${YELLOW}https://golang.org/pkg/${NC}, αφαίρεσέ τα, και τρέξε ${GREEN}go mod tidy${NC}."
    fi
else
    fail "Έλεγχος Standard Library (Δεν βρέθηκε go.mod. Τρέξε 'go mod init <όνομα_root_φακέλου>')"
fi

echo -e "\n${YELLOW}--- 🔍 Έλεγχος στατικών σφαλμάτων και πολυπλοκότητας (DRY, SOLID, SoC) ---${NC}"
if go vet ./...; then
    pass "go vet ./..."
else
    fail "go vet (Βρέθηκαν ύποπτα κομμάτια κώδικα, τρέξε 'go vet ./...' για λεπτομέρειες)"
    echo -e "  👉 ${YELLOW}Οδηγίες Διαχείρισης:${NC}"
    echo -e "  1. Τρέξε χειροκίνητα: ${GREEN}go vet ./...${NC} για να δεις την ακριβή λίστα των σφαλμάτων."
    echo -e "  2. Το go vet συνήθως εντοπίζει λογικά λάθη που ξεφεύγουν από τον compiler."
    echo -e "  3. Συχνά σφάλματα: Λάθος arguments σε Printf, unreachable κώδικας, λάθος αντιγραφή Mutex."
fi

if ! command -v shadow &> /dev/null; then
    fail "shadow δεν βρέθηκε! Τρέξε: go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest"
else
    if go vet -vettool=$(which shadow) ./... >/dev/null 2>&1; then
        pass "go vet -vettool=\$(which shadow) (Έλεγχος Shadowing)"
    else
        fail "shadow (Βρέθηκαν μεταβλητές που επισκιάζουν (shadow) άλλες!)"
        echo -e "  👉 ${YELLOW}Οδηγίες Διαχείρισης:${NC}"
        echo -e "  1. Τρέξε χειροκίνητα: ${GREEN}go vet -vettool=\$(which shadow) ./...${NC} για να δεις λεπτομέρειες."
        echo -e "  2. Συνήθως συμβαίνει με τη χρήση του ${YELLOW}:=${NC} μέσα σε if/for (π.χ. err := ...)."
        echo -e "  3. Χρησιμοποίησε ${YELLOW}=${NC} αν ήθελες να αναθέσεις τιμή στην εξωτερική μεταβλητή, αλλιώς άλλαξε το όνομά της."
    fi
fi

if ! command -v staticcheck &> /dev/null; then
    fail "staticcheck δεν βρέθηκε! Τρέξε: go install honnef.co/go/tools/cmd/staticcheck@latest"
else
    if staticcheck ./... 2>/dev/null; then
        pass "staticcheck ./..."
    else
        fail "staticcheck (Βρέθηκαν θέματα, τρέξε 'staticcheck ./...' για λεπτομέρειες)"
        echo -e "  👉 ${YELLOW}Οδηγίες Διαχείρισης:${NC}"
        echo -e "  1. Τρέξε χειροκίνητα: ${GREEN}staticcheck ./...${NC} για να δεις την ακριβή λίστα των σφαλμάτων."
        echo -e "  2. Διάβασε τον κωδικό του σφάλματος (π.χ. SA1006) και διόρθωσε τον κώδικα."
        echo -e "  3. Συνήθως αφορά αχρησιμοποίητες μεταβλητές, λάθος τύπους ή deprecated συναρτήσεις."
    fi
fi

if ! command -v errcheck &> /dev/null; then
    fail "errcheck δεν βρέθηκε! Τρέξε: go install github.com/kisielk/errcheck@latest"
else
    if errcheck ./... >/dev/null 2>&1; then
        pass "errcheck (Δεν βρέθηκαν unhandled errors)"
    else
        fail "errcheck (Βρέθηκαν unhandled errors!)"
        echo -e "  👉 ${YELLOW}Οδηγίες Διαχείρισης:${NC}"
        echo -e "  1. Τρέξε χειροκίνητα: ${GREEN}errcheck ./...${NC} για να δεις ποια errors αγνοήθηκαν."
        echo -e "  2. Βεβαιώσου ότι κάνεις πάντα 'if err != nil { ... }' ή χρησιμοποιείς '_' αν εσκεμμένα αγνοείς το error."
    fi
fi

if ! command -v nilaway &> /dev/null; then
    fail "nilaway δεν βρέθηκε! Τρέξε: go install go.uber.org/nilaway/cmd/nilaway@latest"
else
    if nilaway ./... >/dev/null 2>&1; then
        pass "nilaway (Δεν βρέθηκαν πιθανά nil pointer panics)"
    else
        fail "nilaway (Βρέθηκαν πιθανά nil pointer panics!)"
        echo -e "  👉 ${YELLOW}Οδηγίες Διαχείρισης:${NC}"
        echo -e "  1. Τρέξε χειροκίνητα: ${GREEN}nilaway ./...${NC} για να δεις λεπτομέρειες."
        echo -e "  2. Έλεγξε τον κώδικα για πιθανά nil pointers πριν γίνει dereference."
    fi
fi

if ! command -v revive &> /dev/null; then
    fail "revive δεν βρέθηκε! Τρέξε: go install github.com/mgechev/revive@latest"
else
    if revive ./... >/dev/null 2>&1; then
        pass "revive (Έλεγχος σχολίων και Style)"
    else
        fail "revive (Λείπουν σχόλια ή βρέθηκαν θέματα ονοματολογίας, τρέξε 'revive ./...' για λεπτομέρειες)"
        echo -e "  👉 ${YELLOW}Οδηγίες Διαχείρισης:${NC}"
        echo -e "  1. Τρέξε χειροκίνητα: ${GREEN}revive ./...${NC} για να δεις τα μηνύματα."
        echo -e "  2. Έλεγξε αν όλες οι δημόσιες (exported) συναρτήσεις/δομές έχουν godoc σχόλιο."
        echo -e "  3. Βεβαιώσου ότι το σχόλιο ξεκινάει με το όνομα της συνάρτησης."
        echo -e "  4. Διόρθωσε τυχόν λάθη στην ονοματολογία (π.χ. underscores αντί για camelCase)."
    fi
fi

if ! command -v gocyclo &> /dev/null; then
    fail "gocyclo δεν βρέθηκε! Τρέξε: go install github.com/fzipp/gocyclo/cmd/gocyclo@latest"
else
    if gocyclo -over 15 . >/dev/null 2>&1; then
        pass "gocyclo (Complexity < 15)"
    else
        fail "gocyclo (Κάποιες συναρτήσεις είναι πολύ πολύπλοκες!)"
        echo -e "\n${YELLOW}--- Top 5 Πιο Πολύπλοκες Συναρτήσεις ---${NC}"
        gocyclo -top 5 .
        echo -e "\n${YELLOW}--- Πίνακας Βαθμολογίας Κυκλωματικής Πολυπλοκότητας ---${NC}"
        echo -e "💡 Η πολυπλοκότητα μετράει τα ανεξάρτητα μονοπάτια (if, for, case) μέσα σε μια συνάρτηση."
        echo -e "  1 - 10: ${GREEN}Απλή${NC} (Εύκολη στην κατανόηση και τον έλεγχο)"
        echo -e " 11 - 20: ${YELLOW}Μέτρια${NC} (Το όριό μας είναι το 15. Προτείνεται refactor)"
        echo -e " 21 - 50: ${RED}Πολύπλοκη${NC} (Υψηλός κίνδυνος bugs)"
        echo -e "   > 50: ${RED}Μη ελέγξιμη${NC} (Απαιτεί άμεσο refactoring)\n"
    fi
fi

if ! command -v misspell &> /dev/null; then
    fail "misspell δεν βρέθηκε! Τρέξε: go install github.com/client9/misspell/cmd/misspell@latest"
else
    # Πρόσθεσε εδώ λέξεις (με κόμμα) που θέλεις να αγνοούνται (π.χ. "word1,word2")
    IGNORE_WORDS="mycustomword"
    if misspell -i "$IGNORE_WORDS" -error . >/dev/null 2>&1; then
        pass "misspell (Δεν βρέθηκαν ορθογραφικά λάθη)"
    else
        fail "misspell (Βρέθηκαν ορθογραφικά λάθη σε αρχεία!)"
        echo -e "  👉 ${YELLOW}Οδηγίες Διαχείρισης:${NC}"
        echo -e "  1. Τρέξε χειροκίνητα: ${GREEN}misspell -i \"$IGNORE_WORDS\" -error .${NC} για να δεις τα λάθη."
        echo -e "  2. Ή άσε το εργαλείο να τα διορθώσει αυτόματα: ${GREEN}misspell -i \"$IGNORE_WORDS\" -w .${NC}"
        echo -e "  3. Για να εξαιρέσεις μια λέξη, πρόσθεσέ την στη μεταβλητή ${YELLOW}IGNORE_WORDS${NC} στο check.sh."
    fi
fi

echo -e "\n${YELLOW}--- 🔒 Έλεγχος Ασφαλείας Κώδικα και Εξωτερικών Κινδύνων (Security) ---${NC}"
if ! command -v gosec &> /dev/null; then
    fail "gosec δεν βρέθηκε! Τρέξε: go install github.com/securego/gosec/v2/cmd/gosec@latest"
else
    if gosec -quiet ./... >/dev/null 2>&1; then
        pass "gosec (Δεν βρέθηκαν κενά ασφαλείας)"
    else
        fail "gosec (Βρέθηκαν πιθανά κενά ασφαλείας!)"
        echo -e "  👉 ${YELLOW}Οδηγίες Διαχείρισης:${NC}"
        echo -e "  1. Τρέξε χειροκίνητα: ${GREEN}gosec ./...${NC} για να δεις την αναλυτική αναφορά."
        echo -e "  2. Διάβασε τον κωδικό του σφάλματος (π.χ. G104, G110) και διόρθωσε την ευπάθεια."
        echo -e "  3. Αν πρόκειται για false positive, πρόσθεσε το σχόλιο ${YELLOW}// #nosec GXXX${NC} (όπου GXXX ο κωδικός) ακριβώς πάνω από τη γραμμή."
    fi
fi

if ! command -v govulncheck &> /dev/null; then
    fail "govulncheck δεν βρέθηκε! Τρέξε: go install golang.org/x/vuln/cmd/govulncheck@latest"
else
    if govulncheck ./... >/dev/null 2>&1; then
        pass "govulncheck (Δεν βρέθηκαν γνωστές ευπάθειες (vulnerabilities) σε εξαρτήσεις (Dependencies) ή στην ίδια την Go!)"
    else
        fail "govulncheck (Βρέθηκαν γνωστές ευπάθειες (CVEs)!)"
        echo -e "  👉 ${YELLOW}Οδηγίες Διαχείρισης:${NC}"
        echo -e "  1. Τρέξε χειροκίνητα: ${GREEN}govulncheck ./...${NC} για να δεις ποιες βιβλιοθήκες είναι ευάλωτες."
        echo -e "  2. Αναβάθμισε τις ευάλωτες βιβλιοθήκες τρέχοντας π.χ. ${GREEN}go get github.com/πακέτο@latest${NC}."
        echo -e "  3. Αν η ευπάθεια αφορά την ίδια την Go (Standard Library), εξέτασε το ενδεχόμενο αναβάθμισης της έκδοσης της Go."
    fi
fi

echo -e "\n${YELLOW}--- ⚡ Βελτιστοποίηση Μνήμης και Δομών (Mechanical Sympathy) ---${NC}"
if ! command -v fieldalignment &> /dev/null; then
    fail "fieldalignment δεν βρέθηκε! Τρέξε: go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest"
else
    if fieldalignment ./... >/dev/null 2>&1; then
        pass "fieldalignment (Βελτιστοποίηση μνήμης Structs)"
    else
        fail "fieldalignment (Βρέθηκαν structs με μη βέλτιστη διάταξη πεδίων!)"
        echo -e "  👉 ${YELLOW}Οδηγίες Διαχείρισης:${NC}"
        echo -e "  1. Τρέξε χειροκίνητα: ${GREEN}fieldalignment ./...${NC} για να δεις ποια structs σπαταλάνε μνήμη (padding)."
        echo -e "  2. Αναδιάταξε τα πεδία από το μεγαλύτερο στο μικρότερο (π.χ. []byte, int64, int32, bool)."
        echo -e "  3. Εναλλακτικά, άφησε το εργαλείο να τα φτιάξει αυτόματα τρέχοντας: ${GREEN}fieldalignment -fix ./...${NC}"
    fi
fi

if go test -bench=. -benchmem ./... >/dev/null 2>&1; then
    pass "go test -bench (Έλεγχος Ταχύτητας & Memory Allocations)"
else
    fail "go test -bench (Τα Benchmarks απέτυχαν!)"
    echo -e "  👉 ${YELLOW}Οδηγίες Διαχείρισης:${NC}"
    echo -e "  1. Τρέξε χειροκίνητα: ${GREEN}go test -bench=. -benchmem ./...${NC} για να δεις τα αποτελέσματα."
    echo -e "  2. Έλεγξε τη στήλη ${YELLOW}allocs/op${NC}. Αν κάνεις περιττές δεσμεύσεις μνήμης, δοκίμασε προ-δέσμευση χωρητικότητας (capacity) στα slices/maps."
    echo -e "  3. Έλεγξε τη στήλη ${YELLOW}ns/op${NC} για να δεις τον χρόνο εκτέλεσης της κάθε λούπας."
    echo -e "  4. Βεβαιώσου ότι δεν υπάρχουν errors ή panics μέσα στις ίδιες τις συναρτήσεις Benchmark."
fi

echo -e "\n${YELLOW}--- 🧪 Εκτέλεση δοκιμών, ελέγχων απόδοσης και κάλυψης (TDD) ---${NC}"
if ! command -v gcc &> /dev/null; then
    fail "go test -race (Λείπει το GCC. Εγκατάστησε το MinGW-w64 για να τρέξει ο Race Detector!)"
else
    if CGO_ENABLED=1 go test -race ./... 2>/dev/null; then
        pass "go test -race"
    else
        fail "go test -race (Βρέθηκαν Race Conditions!)"
        echo -e "  👉 ${YELLOW}Οδηγίες Διαχείρισης:${NC}"
        echo -e "  1. Τρέξε χειροκίνητα: ${GREEN}CGO_ENABLED=1 go test -race ./...${NC} για να δεις την αναλυτική αναφορά."
        echo -e "  2. Ψάξε για το ${RED}WARNING: DATA RACE${NC} στο τερματικό."
        echo -e "  3. Δες ποιες ακριβώς γραμμές κώδικα (Read at... / Previous write at...) συγκρούονται."
        echo -e "  4. Χρησιμοποίησε ${YELLOW}sync.Mutex${NC} ή ${YELLOW}Channels${NC} για να συγχρονίσεις την πρόσβαση στα δεδομένα."
    fi
fi

if go test -fuzz=Fuzz -fuzztime=10s ./... >/dev/null 2>&1; then
    pass "go test -fuzz (10s fuzzing)"
else
    fail "go test -fuzz (Βρέθηκαν σφάλματα με τυχαία δεδομένα!)"
    echo -e "  👉 ${YELLOW}Οδηγίες Διαχείρισης:${NC}"
    echo -e "  1. Τρέξε χειροκίνητα: ${GREEN}go test -fuzz=Fuzz -fuzztime=10s ./...${NC} για να δεις το error."
    echo -e "  2. Δες το \"κακό\" input που αποθηκεύτηκε αυτόματα στον φάκελο ${YELLOW}testdata/fuzz/${NC}."
    echo -e "  3. Κάνε debug τρέχοντας απλά ${GREEN}go test ./...${NC} (το failing input θα τρέχει πλέον ως κανονικό test)."
fi

if COVER_OUT=$(go test -coverprofile=coverage.out -coverpkg=./... ./... 2>&1); then
    if [ -f coverage.out ]; then
        # Αφαίρεση εξαιρέσεων από το coverage (π.χ. Mocks, Generated code)
        IGNORE_COVERAGE_PATTERNS=(
            "mock_"
            "/mocks/"
            "_mock.go"
        )
        for pattern in "${IGNORE_COVERAGE_PATTERNS[@]}"; do
            grep -v "$pattern" coverage.out > coverage_temp.out && mv coverage_temp.out coverage.out
        done

        # Υπολογισμός συνολικού ποσοστού κάλυψης
        COVERAGE=$(go tool cover -func=coverage.out | grep '^total:' | awk '{print $3}' | tr -d '%')
        THRESHOLD=80.0
        
        if awk "BEGIN {exit !($COVERAGE >= $THRESHOLD)}"; then
            pass "go test coverage ($COVERAGE% >= $THRESHOLD%)"
        else
            fail "go test coverage ($COVERAGE% - Απαιτείται τουλάχιστον $THRESHOLD%)"
            echo -e "  👉 ${YELLOW}Οδηγίες Διαχείρισης:${NC}"
            echo -e "  1. Το Test Coverage είναι κάτω από το επιτρεπτό όριο!"
            echo -e "  2. Άνοιξε το HTML Report στον browser για να δεις ποιες γραμμές δεν έχουν ελεγχθεί..."
            go tool cover -html=coverage.out
            echo -e "  3. Γράψε επιπλέον tests (π.χ. table-driven) για να καλύψεις τις κόκκινες γραμμές."
        fi
    else
        fail "go test coverage (Δεν βρέθηκαν tests!)"
    fi
else
    fail "go test coverage (Τα tests απέτυχαν!)"
    echo -e "${YELLOW}Λεπτομέρειες Σφάλματος:${NC}\n$COVER_OUT"
    echo -e "  👉 ${YELLOW}Οδηγίες Διαχείρισης:${NC}"
    echo -e "  1. Εντόπισε το test που απέτυχε από τα παραπάνω μηνύματα."
    echo -e "  2. Ανοίγουμε το HTML Report στον browser (εφόσον δημιουργήθηκε) για να δεις τι εκτελέστηκε..."
    if [ -f coverage.out ]; then go tool cover -html=coverage.out; fi
    echo -e "  3. Μπορείς πάντα να το δεις χειροκίνητα με: ${GREEN}go tool cover -html=coverage.out${NC}"
fi

echo -e "\n${YELLOW}--- ⏱️ Ταχύτητα Εκτέλεσης (Time Limitation - 5 Minutes Limit) ---${NC}"
if go build -o temp_bin . ; then
    pass "go build (Επιτυχής μεταγλώττιση για δοκιμές χρόνου)"
    
    SAMPLES_DIR="samples"
    if [ -d "$SAMPLES_DIR" ]; then
        echo -e "  Εκτέλεση αρχείων από τον φάκελο '${SAMPLES_DIR}' (Όριο: 5 λεπτά / 300s)..."
        for file in "$SAMPLES_DIR"/*; do
            [ -e "$file" ] || continue # Παράλειψη αν ο φάκελος είναι άδειος
            
            start=$SECONDS
            # Η εντολή timeout διακόπτει την εκτέλεση αν ξεπεράσει τα 300 δευτερόλεπτα (5 λεπτά)
            timeout 300s ./temp_bin "$file" > /dev/null 2>&1
            exit_status=$?
            duration=$(( SECONDS - start ))
            
            if [ $exit_status -eq 124 ]; then
                fail "Εκτέλεση: $file (Time Limit Exceeded! Το πρόγραμμα ξεπέρασε τα 300s)"
            else
                pass "Εκτέλεση: $file (Ολοκληρώθηκε σε ${duration}s)"
            fi
        done
    else
        echo -e "  👉 ${YELLOW}Info: Δεν βρέθηκε φάκελος '${SAMPLES_DIR}'. Παράλειψη ελέγχου ταχύτητας εκτέλεσης.${NC}"
    fi
    rm -f temp_bin
else
    fail "go build (Αποτυχία Μεταγλώττισης)"
fi

echo -e "\n${YELLOW}--- Summary ---${NC}"
echo -e "Passed: ${GREEN}${PASS_COUNT}${NC}"
echo -e "Failed: ${RED}${FAIL_COUNT}${NC}"

if [ "$FAIL_COUNT" -eq 0 ]; then
    echo -e "\n${GREEN}✅ Όλοι οι έλεγχοι ολοκληρώθηκαν με επιτυχία! Έτοιμος για Review.${NC}"
    exit 0
else
    echo -e "\n${RED}❌ Προσοχή: $FAIL_COUNT έλεγχοι απέτυχαν. Διόρθωσε τα σφάλματα και δοκίμασε ξανά.${NC}"
    exit 1
fi