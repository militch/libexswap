#include "exchange.h"
#include "token.h"

#include <stdio.h>
#include <stdlib.h>


amm_exchange_t *amm_exchange_new(){
    amm_exchange_t *ex = malloc(sizeof(amm_exchange_t));
    if (!ex) {
        return NULL;
    }
    return ex;
}


/**
 * 增加流动性
 * @param ex 交易所上下文 
 * @param min_liquidity 最小流动性
 * @param max_tokens 最大通证数量
 * @param value Delta 数量
 */
int amm_add_liquidity(amm_exchange_t *ex, unsigned int min_liquidity,
                      unsigned int max_tokens, unsigned int value) {
    if (max_tokens <= 0 || value <= 0) {
        return 1;
    }
    int total_liquidity = ex->total_supply;
    int token_balance = ex->token_balance;
    int ex_balance = ex->balance;
    if (total_liquidity > 0) {
        if (min_liquidity <= 0) {
            return 1;
        }
        int delta_reserve = ex_balance - value;
        int token_reserve = token_balance;
        int token_amount = 0;
        int liquidity_minted = 0; 
        if (delta_reserve > 0) {
            token_amount = value * token_reserve / delta_reserve + 1;
            liquidity_minted = value * total_liquidity / delta_reserve;
        }
        if (max_tokens < token_amount || liquidity_minted < min_liquidity) {
            printf("max_tokens: %d, token_amount=%d\n", max_tokens, token_amount);
            printf("min_liquidity: %d, liquidity_minted=%d\n", min_liquidity, liquidity_minted);
            return 1;
        }
        printf("result: token_amount=%08d, liquidity=%08d\n", token_amount,  liquidity_minted);
        ex->token_balance += token_amount;
        ex->total_supply = total_liquidity + liquidity_minted;
        return 0;
    } else {
        int token_amount = max_tokens;
        int initial_liquidity = value;
        printf("result: token_amount=%08d, liquidity=%08d\n", token_amount,  initial_liquidity);
        ex->token_balance += token_amount;
        ex->total_supply = initial_liquidity;
        return 0;
    }
}

// 获取输入价格
unsigned int amm_get_input_price(amm_exchange_t *ex, unsigned int input_amount,
                        unsigned int input_reserve, unsigned int output_reserve){
    if (input_reserve <= 0 || output_reserve <= 0) {
        return 0;
    }
    unsigned int input_amount_with_fee = input_amount * 10;
    unsigned int numerator = input_amount_with_fee * output_reserve;
    unsigned int denominator = (input_reserve * 1000) + input_amount_with_fee;
    unsigned int a = numerator / denominator;
    return a;
}

unsigned int amm_get_delta_to_token_input_price(amm_exchange_t *ex, unsigned int delta_sold) {
    if (delta_sold <= 0) {
        return 0;
    }
    unsigned int delta_reserve = ex->balance - delta_sold;
    unsigned int token_reserve = ex->token_balance;
    return amm_get_input_price(ex, delta_sold, delta_reserve, token_reserve);
}

int amm_delta_to_token_input_price(amm_exchange_t *ex, unsigned int delta_sold, unsigned int min_tokens){
    unsigned int price = amm_get_delta_to_token_input_price(ex, delta_sold);
    if (price < min_tokens) {
        return 1;
    }
    ex->token_balance -= price;
    return 0;
}

int amm_get_output_price(amm_exchange_t *ex, int output_amount,
                        int input_reserve, int output_reserve){

    return 0;
}

/**
 * 设置函数
 */
int amm_set_token_balance_getfn(amm_exchange_t* ex, int (*fn)()){
    return 0;
}
