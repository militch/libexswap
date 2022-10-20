#include <stdio.h>
#include <stdlib.h>

#include "exchange.h"

int main(int argc, char* argv[]){
    // 通证
    amm_token_t *token;
    token = amm_token_new(1);


    // 交易所
    amm_exchange_t *ex;
    ex = amm_exchange_new(token);
    amm_add_liquidity(ex, 1, 1, 1);
    //int n = tokena->balance;
    //printf("n=%s\n", tokena);
    return 0;
}
