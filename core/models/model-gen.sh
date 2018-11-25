#! /bin/bash

# provide argument to this script like so: "file store"

snakeCase=""
camelCase=""
variableCase=""
receiver=""
allCapped=""

count=0
for a in $1; do
    receiver="${receiver}${a:0:1}"

    capped=$(tr '[:lower:]' '[:upper:]' <<< ${a:0:1})${a:1}
    camelCase="${camelCase}${capped}"
    
    if [ $count -eq 0 ]; then
        echo "on first iteration of loop"
        snakeCase=$a
        variableCase=$a
        allCapped=${a^^}
    else
        snakeCase="${snakeCase}_$a"
        variableCase="${variableCase}${capped}"
        allCapped="${allCapped}_${a^^}"
    fi
   
    # camelCase="${camelCase}$(tr '[:lower:]' '[:upper:]' <<< ${a:0:1})${a:1}"
    (( count++ ))
done

echo "snakeCase: $snakeCase"
echo "camelCase: $camelCase"
echo "variableCase: $variableCase"
echo "receiver: $receiver"
echo "allCapped: $allCapped"

echo "Are these values ok? (y/N)"
read proceed
if [ $proceed -a $proceed != 'y' ]; then 
    echo 'not given permission to proceed, exiting...'
    exit 0
fi

echo "generating files..."
# exit 1 # remove this statement when done testing

# model.go
echo "package models

type $camelCase struct { 
	// DB-mapped 
	ID            int    \`json:\"id\"\` 
}" > "$snakeCase.go"

# model_store.go
echo "package models

// ${camelCase}Store is the interface that all storers of ${camelCase} must implement
type ${camelCase}Store interface {
    GetByID(int) (*$camelCase, error)
    Save(*$camelCase) error
    Delete(int) error
}
" > "${snakeCase}_store.go"

# model_pg_store.go
pgStoreReceiver="${receiver}ps"
echo "package models

import (
    \"os\"

    \"github.com/briand787b/rfs/core/postgres\"

    \"github.com/jmoiron/sqlx\"
    \"github.com/pkg/errors\"
)

type ${variableCase}PGStore struct {
    db postgres.ExtFull
}

// New${camelCase}PGStore returns a ${camelCase}Store backed by Postgresql
func New${camelCase}PGStore(db *sqlx.DB) ${camelCase}Store { 
    return &${variableCase}PGStore{db: postgres.GetExtFull(os.Stdout)} 
}

func ($pgStoreReceiver *${variableCase}PGStore) GetByID(id int) (*${camelCase}, error) {
	var ${receiver}Rec $camelCase
    if err := sqlx.Get(${pgStoreReceiver}.db, &${receiver}Rec, \`
		SELECT
			id AS ID,
            -- FILL IN THE REST OF THIS\!\!\!
		FROM
			${snakeCase}s
		WHERE
			id = \$1;\`,
		id,
	); err != nil {
		return nil, errors.Wrap(err, \"failed to execute query\")
	}

	return &${receiver}Rec, nil
}

func ($pgStoreReceiver *${variableCase}PGStore) Save($receiver *$camelCase) error {
	var saveID int
	if err := sqlx.Get($pgStoreReceiver.db, &saveID, \`
		INSERT INTO ${snakeCase}s
		(
			-- FILL IN THE REST OF THIS\!\!\!
		)
		VALUES
		(
			-- FILL IN THE REST OF THIS\!\!\!
		)
		RETURNING id;\`,
	    // FILL IN THE REST OF THIS\!\!\!
	); err != nil {
		err = errors.Wrap(err, \"failed to execute query\")
	}

    $receiver.ID = saveID
	return nil
}

func ($pgStoreReceiver *${variableCase}PGStore) Delete(id int) error {
	var delID int
    if err := sqlx.Get($pgStoreReceiver.db, &delID, \`
		DELETE FROM ${snakeCase}s
		WHERE
			id = \$1;\`,
		id,
	); err != nil {
		return errors.Wrap(err, \"could not execute query\")
	}

    if delID == 0 {
		return errors.New(\"row was not actually deleted\")
	}

	return nil
}
" > "${snakeCase}_pg_store.go"

# model_pg_store_test.go
echo "package models

import (
	\"fmt\"
	\"reflect\"
	\"testing\"
)

// TODO: 
//  1) Inside of test_util_test.go create
//      a) \`const  $allCapped\`
//      b) \`var ${variableCase}Store *${variableCase}PGStore\`
//      c) \`${variableCase}Store = New${camelCase}PGStore(db).(*${variableCase}PGStore)\`
//          in \`SetUpDepencencies()\`

// Create$camelCase creates a $camelCase.  If the passed
// $camelCase pointer is nil, then the default $camelCase is
// created.  Otherwise, the provided $camelCase is created
// in the database.  The returned $camelCase pointer should
// only be used if the provided $camelCase pointer is nil.
func Create$camelCase(ts *TestStruct, $receiver *$camelCase) *$camelCase {
	if $receiver == nil {
		$receiver = &$camelCase{
            // FILL THIS IN WITH CORRECT FIELDS!!!
        }
	}

	if err := ${variableCase}Store.Save($receiver); err != nil {
		defer (*ts.CF)()
		ts.Fatal(\"could not create media: \", err)
	}

	fmt.Println(\"created $snakeCase with ID: \", $receiver.ID)

	ts.CF.Add(func() {
		if err := ${variableCase}Store.Delete($receiver.ID); err != nil {
			ts.Fatal(\"could not delete created $snakeCase: \", err)
		}
	})

	ts.ParentIDs[$allCapped] = $receiver.ID

	return $receiver
}

func TestSaveGetByIDDelete$camelCase(t *testing.T) {
	$receiver := &$camelCase{
		// FILL THIS IN WITH CORRECT FIELDS!!!
	}

	if ${variableCase}Store == nil {
		t.Fatal(\"$camelCase STORE IS NIL\")
	}

	if err := ${variableCase}Store.Save($receiver); err != nil {
		t.Fatal(err)
	}

	ret${receiver^^}, err := ${variableCase}Store.GetByID($receiver.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(ret${receiver^^}, $receiver) {
		t.Fatalf(\"returned $camelCase %+v does not equal saved $snakeCase %+v\n\",
			ret${receiver^^}, $receiver,
		)
	}

	if err := ${variableCase}Store.Delete($receiver.ID); err != nil {
		t.Fatal(err)
	}

	if _, err := ${variableCase}Store.GetByID($receiver.ID); err == nil {
		t.Fatalf(\"deleted $snakeCase is still found by id %v\n\", $receiver.ID)
	}
}" > "${snakeCase}_pg_store_test.go"

# ./modeltest/model_store_mock.go
echo "package modeltest

import \"github.com/briand787b/rfs/core/models\"

// ${camelCase}StoreMock is a mocked implementation of ${camelCase}Store
type ${camelCase}StoreMock struct {
	// function return info
	GetByID${receiver^^} *models.$camelCase
	GetByIDE error
	Save${receiver^^}    *models.$camelCase
	SaveE    error
	Delete${receiver^^}  *models.$camelCase
	DeleteE  error

	// Spy fields
	CallCount int
}

// New${camelCase}StoreMockGetByID returns a ${camelCase}Store ready to mock the GetByID method
func New${camelCase}StoreMockGetByID($receiver *models.$camelCase, e error) *${camelCase}StoreMock {
	return &${camelCase}StoreMock{
		GetByID${receiver^^}: $receiver,
		GetByIDE: e,
	}
}

// New${camelCase}StoreMockSave returns a ${camelCase}Store ready to mock the Save method
func New${camelCase}StoreMockSave($receiver *models.$camelCase, e error) *${camelCase}StoreMock {
	return &${camelCase}StoreMock{
		Save${receiver^^}: $receiver,
		SaveE: e,
	}
}

// New${camelCase}StoreMockDelete returns a ${camelCase}Store ready to mock the Delete method
func New${camelCase}StoreMockDelete($receiver *models.$camelCase, e error) *${camelCase}StoreMock {
	return &${camelCase}StoreMock{
		Delete${receiver^^}: $receiver,
		DeleteE: e,
	}
}

// GetByID returns the ${camelCase}StoreMock's GetByID${receiver^^} and GetByIDE
func (${receiver}sm *${camelCase}StoreMock) GetByID(id int) (*models.$camelCase, error) {
	${receiver}sm.CallCount++
	return ${receiver}sm.GetByID${receiver^^}, ${receiver}sm.GetByIDE
}

// Save returns the ${camelCase}StoreMock's Save${receiver^^} and SaveE
func (${receiver}sm *${camelCase}StoreMock) Save(id int) (*models.$camelCase, error) {
	${receiver}sm.CallCount++
	return ${receiver}sm.Save${receiver^^}, ${receiver}sm.SaveE
}

// Delete returns the ${camelCase}StoreMock's Delete${receiver^^} and DeleteE
func (${receiver}sm *${camelCase}StoreMock) Delete(id int) (*models.$camelCase, error) {
	${receiver}sm.CallCount++
	return ${receiver}sm.Delete${receiver^^}, ${receiver}sm.DeleteE
}" > "./modeltest/${snakeCase}_store_mock.go"

echo "done"