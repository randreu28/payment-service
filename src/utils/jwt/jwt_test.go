package jwt

import "testing"

func TestJwt(t *testing.T) {
	payload := Payload{
		Id:    1,
		Owner: "test",
	}

	jwt, err := CreateJwt(payload.Id, payload.Owner)
	if err != nil {
		t.Fatal(err)
	}

	token, err := ParseToken(jwt)
	if err != nil {
		t.Fatal(err)
	}

	if token.Payload != payload {
		t.Fatal("Payloads not equal")
	}
}
