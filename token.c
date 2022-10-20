#include "token.h"

#include <stdlib.h>


amm_token_t *amm_token_new(int total_supply) {
    amm_token_t *token = malloc(sizeof(amm_token_t));
    if (!token) {
        return NULL;
    }
    return token;
}

void amm_token_mint(amm_token_t *t, int a){

}


