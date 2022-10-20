#include "exchange.h"
#include "token.h"

#include <stdio.h>
#include <stdlib.h>

amm_exchange_t *amm_exchange_new(amm_token_t *t){
    amm_exchange_t *ex = malloc(sizeof(amm_exchange_t));
    if (!ex) {
        return NULL;
    }
    ex->token = t;
    return ex;
}


/**
 * 增加流动性
 * @param ex 交易所上下文 
 * @param min_liquidity 
 */
int amm_add_liquidity(amm_exchange_t *ex, int min_liquidity,
                      int max_tokens, int value) {

    int total_liquidity = ex->total_supply;
    int token_balance = ex->token_balance;
    int ex_balance = ex->balance;
    printf("total_liquidity: %d\n", total_liquidity);
    printf("ex_balance: %d\n", ex_balance);
    printf("token_balance: %d\n", token_balance);
    if (total_liquidity > 0) {
        if (min_liquidity <= 0){
            return 1;
        }
        int delta_reserve = ex_balance - value;
        int token_reserve = token_balance;
        int token_amount = value * token_reserve / delta_reserve + 1;
        int liquidity_minted = value * total_liquidity / delta_reserve;
        if (max_tokens < token_amount || liquidity_minted < min_liquidity) {
            return 1;
        }
        ex->total_supply =  total_liquidity + liquidity_minted;
    } else {
        int token_amount = max_tokens;
        int initial_liquidity = ex_balance;
        ex->total_supply = initial_liquidity;
    }
    return 0;
    //int t = ex->total_supply;
    //ex->total_supply = 990;
}

