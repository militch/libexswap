#ifndef EXCHANGE_H_
#define EXCHANGE_H_

#include "token.h"

struct amm_exchange {
    int total_supply;
    int token_balance;
    int balance;
    amm_token_t *token;
};

#define amm_exchange_t struct amm_exchange

// 构造交易所
amm_exchange_t *amm_exchange_new(amm_token_t*);
// 添加流动性
int amm_add_liquidity(amm_exchange_t*, int, int, int);

#endif // EXCHANGE_H_
