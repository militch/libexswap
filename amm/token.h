#ifndef TOKEN_H_
#define TOKEN_H_

struct amm_token {
    int total_supply;
};

#define amm_token_t struct amm_token

// 构造通证
amm_token_t *amm_token_new(int);

// 铸造通证
void amm_token_mint(amm_token_t*, int);

// 获取余额
int amm_token_get_balance(int);

#endif
