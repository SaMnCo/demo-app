package dbclient

import (
        "github.com/boltdb/bolt"
        "log"
        "strconv"
        "github.com/validatepolicy/model"
        "encoding/json"
        "fmt"
)


type IBoltClient interface {
        OpenBoltDb()
        QueryPolicy(policyNumber string) (model.Policy, error)
        Seed()
}

// Real implementation
type BoltClient struct {
        boltDB *bolt.DB
}

func (bc *BoltClient) OpenBoltDb() {
        var err error
        bc.boltDB, err = bolt.Open("policies.db", 0600, nil)
        if err != nil {
                log.Fatal(err)
        }
}

func (bc *BoltClient) QueryPolicy(policyNumber string) (model.Policy, error) {
        // Allocate an empty Policy instance we'll let json.Unmarhal populate for us in a bit.
        policy := model.Policy{}

        // Read an object from the bucket using boltDB.View
        err := bc.boltDB.View(func(tx *bolt.Tx) error {
                // Read the bucket from the DB
                b := tx.Bucket([]byte("PolicyBucket"))

                // Read the value identified by our policyNumber supplied as []byte
                policyBytes := b.Get([]byte(policyNumber))
                if policyBytes == nil {
                        fmt.Errorf("No policy found for " + policyNumber)
                }
                // Unmarshal the returned bytes into the policy struct we created at
                // the top of the function
                json.Unmarshal(policyBytes, &policy)

                // Return nil to indicate nothing went wrong, e.g no error
                return nil
        })
        // If there were an error, return the error
        if err != nil {
                return model.Policy{}, err
        }
        // Return the Policy struct and nil as error.
        return policy, nil
}


// Start seeding Policies
func (bc *BoltClient) Seed() {
        bc.initializeBucket()
        bc.seedPolicies()
}

// Creates an "PolicyBucket" in our BoltDB. It will overwrite any existing bucket of the same name.
func (bc *BoltClient) initializeBucket() {
        bc.boltDB.Update(func(tx *bolt.Tx) error {
                _, err := tx.CreateBucket([]byte("PolicyBucket"))
                if err != nil {
                        return fmt.Errorf("create bucket failed: %s", err)
                }
                return nil
        })
}

// Seed (n) make-believe Policy objects into the AcountBucket bucket.
func (bc *BoltClient) seedPolicies() {

        total := 10
        for i := 0; i < total; i++ {

                // Generate a key 10000 or larger
                key := strconv.Itoa(100000 + i)

                // Create an instance of our Policy struct
                policy := model.Policy{
                        Id: "DUHSS"+key,
                        Name: "Person_" + strconv.Itoa(i),
                        RiskAddress: strconv.Itoa(i) + " Latitue B5 6AE",
                        ServiceVersion: "1.0",
                }

                // Serialize the struct to JSON
                jsonBytes, _ := json.Marshal(policy)

                // Write the data to the PolicyBucket
                bc.boltDB.Update(func(tx *bolt.Tx) error {
                        b := tx.Bucket([]byte("PolicyBucket"))
                        err := b.Put([]byte(key), jsonBytes)
                        return err
                })
        }
        fmt.Printf("Seeded %v dummy polices...\n", total)
}

