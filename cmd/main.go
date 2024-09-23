package main

import (
	"github.com/houg/go-oauth2-resource/resource"
	"log"
)

func main() {

	// This is a local Keycloak JWK Set endpoint for the master realm.
	jwksURL := "https://impre.zdxlz.com/seal/oauth2/jwks"

	// Get a JWT to parse.
	jwtB64 := "eyJraWQiOiJwRVBhZXZNWHVMNUtxUkMiLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiIxODI4NzYyNzkxODk3NTg3NzEzIiwiYXVkIjoiOWU5YzI5YzM1NWM2NDZhY2FhZiIsIm5iZiI6MTcyNzA3NjkwMywic2NvcGUiOlsib3BlbmlkIiwicHJvZmlsZSJdLCJpc3MiOiJodHRwOi8vaW1wcmUuemR4bHouY29tL3NlYWwiLCJleHAiOjE3MjcwODQxMDMsImlhdCI6MTcyNzA3NjkwMywianRpIjoiMzg3ZmM2YTEtN2ZiMC00NWE1LWJiNGItOGY3MWY0MDcwMGE5In0.GZG6O8Jaxr7lyaDr4yUwaX3XoWGiLD6cXeHAJ9kwfoZOB3ZaXJotQzCBi8uNbOiw718tDyNzkF4o2j5cS8BN9gWCwPNIWe-mjLf9ctL_a068JkLPmEfgikS9-dCFMPIweObrnf0daCLTfktNKXO8hizU6gpiVfUoh92NNyISsi56XjrVBHxo6IlXDU5-gOYXDPw_Ux69O5XOBosbjARDFQVSzsooujg7PeGpmTalPGJoMRnx7Q0-Z8sYIUB7mly9hB8qJ7LLMTuTafrcBrWxmGhcjzDlxEF8nFwATLR9ZOV1yilGlvCZMdpvBHOdwg8_hZ1IjQORbwj3rJYmdQYQIw"

	config := resource.NewConfig(jwksURL, true)
	tokenService := resource.NewTokenServ(config)

	accessToken, err := tokenService.ParseAccessToken(jwtB64)
	if err != nil {
		log.Printf("Error: %s \n", err)
		return
	}
	log.Printf("HasScopes: %v \n", accessToken.HasScopes("openid"))
	log.Printf("HasGrantType: %v \n", accessToken.HasGrantType("authorization_code"))
}
