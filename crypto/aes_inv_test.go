package crypto

import (
	"fmt"
	"log"
	"testing"
)

func Test_AES_INV(t *testing.T) {
	key := []byte("XCY03LX06ZLQN30J") // 16字节的秘钥
	iv := []byte("KM97SH196CXCY6C9")  // 16字节的偏移量

	//plaintext := "Hello, AES!"
	//
	//ciphertext, err := EncryptByinv(key, iv, plaintext)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("Ciphertext:", ciphertext)

	ss := `DFR47CZYbGbgLs1jCFZWLRK3YdE4iSB1lOoMCpsaim3CpfwluMfTkNnWVFIU9lOOGrNZI/ZhQFqcWnzmI5eGAefhTKHuxQwlmBoE36RqEELXqHqwxbXPlHCV0RUkTtoPj1BU+mLu17PVhpQxXJ14JHS/jFdkj8dLVfXtZNrMC2syOIcESiu7niNYmOz3DeKTJpsPuIfQHJvtp2HkL/NESJp3aOYwhJvho0Qnzm4d67Q6me4hyobkP+XmNLW5QgLIOpUEST6ys6F7CdJhuLwIZuiKB+8H5IvuT7B0efP32k0qM4jOu1fywEzu7jrpV4Nt4Aos4CFaGBSg3E6uDfIRj/IBMBFpeurnLwEVfULpLMXNbuP3qz9UMVefEreAfvGpvA1u7zfHDkR82JgeR6S8smjV2KcPYFneDFavBr7w1m3uM8lo1PhEZcspHF8qKRhTO3Vqn2c38Z5SE4gFnIF7DhPQTmsw5XrzA3hK478DpgWvlm+LkVW8N+sOyCK6OnFWsPUT5Xdf0N9ol3Osfak5APOlFfUMHNgwQ9mm1OVbUIletxoTzyjm2wkDRhkct61GyTf/q1zwH+SxqKkgqUTRQNXhHFcmoUcopiatJ4gtSf8nR1v2FdonHqd2v8XMWvawJRD/j/y86Qdx9KispfAS3+3yn1MX22YVNcEySZelmF6s3nff+sv70kZM0rjk1uEqJJOorZzRKhm4PeNisK8j2OSDBERvLBQc7sAb2D57x6CEssPCNaPqSqPQyoOxodHr0zE6lQj1BL77uEogL3rJctPNjuj/8Me5aAJjvCyDycphcwNWfBEp/vvl+ftVUbs2aGlsa2UvjHvxHSkZCUsqOINDvh2smAIZ9HDc47oHkCSwpmhitDud2Vy9QE6schIMy7rw2WGLHe1Ef+ZdE3QV4pqJFXHu27paGxqxq9nPy1aKjIwG9++ClJGit+eJM9Vf1py0QAyzDT4794Yt7SJBOCb7IUoqDqTSxaf1LmlgGb0vsUE1hDoSxu3p6dWkYLWMLqgyjT/nK1Kd+JHpNtAiRZCRmuWoD85uOzCsE4aBW6886DMFQR0KhEy9W0HrnjDT`

	decryptedText, err := DecryptByinv(key, iv, ss)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Decrypted text:", decryptedText)
}
