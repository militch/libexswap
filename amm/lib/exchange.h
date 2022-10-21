#ifndef EXCHANGE_H_
#define EXCHANGE_H_

#include "token.h"


/**
 * total_supply 总供给量（流动性）
 * token_balance 通证余额
 * balance 通证余额 
 */
struct amm_exchange {
    unsigned int total_supply;
    unsigned int balance;
    unsigned int token_balance;
};

#define amm_exchange_t struct amm_exchange

/**
 * 构造交易所
 * @param total_supply 总供给量
 * @param balance 通证余额
 * @param token_balance 通证余额
 */
amm_exchange_t *amm_exchange_new(unsigned int, unsigned int, unsigned int);

/**
 * 增加流动性
 * @param ex 交易所上下文 
 * @param min_liquidity 最小流动性
 * @param max_tokens 最大通证数量
 * @param value 数量
 */
int amm_add_liquidity(amm_exchange_t*, unsigned int, unsigned int, unsigned int);

/**
 * 移除流动性
 * @param ex 交易所上下文 
 * @param min_liquidity 最小流动性
 * @param max_tokens 最大通证数量
 * @param value 数量
 */
int amm_remove_liquidity(amm_exchange_t*, int, int, int);

/**
 * 获取输入价格
 * @param ex 交易所上下文
 * @param input_amount 输入数量
 * @param input_reserve 输入储备金
 * @param output_reserve 输出储备金
 */
unsigned int amm_get_input_price(amm_exchange_t*, unsigned int, unsigned int, unsigned int);

unsigned int amm_get_delta_to_token_input_price(amm_exchange_t*, unsigned int);

int amm_delta_to_token_input_price(amm_exchange_t*, unsigned int, unsigned int);

/**
 * 获取输出价格
 * @param ex 交易所上下文
 * @param output_amount 输出数量
 * @param input_reserve 输入储备金
 * @param output_reserve 输出储备金
 */
int amm_get_output_price(amm_exchange_t*, int, int, int);

#endif
