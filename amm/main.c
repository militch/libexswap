#include <stdio.h>
#include <stdlib.h>

#include "exchange.h"

int getfn(){
    printf("abc\n"); 
    return 0;
}

unsigned int yuan2fen(unsigned int n){
    return n*100;
}

unsigned int fen2yuan(unsigned int n){
    return n/100;
}

int main(int argc, char* argv[]){
    
    // 通证
    // amm_token_t *token;
    // token = amm_token_new(1);

    amm_exchange_t *ex;
    // 初始化交换池
    ex = amm_exchange_new(0, 0, 0);
    printf("-- amm_add_liquidity: 00\n");
    ex->balance += yuan2fen(1000);
    int err = amm_add_liquidity(ex, 0, yuan2fen(10000), yuan2fen(100));
    if (err) {
        printf("amm_add_liquidity() = false\n");
        return 1;
    }

    // printf("-- amm_get_input_price:\n");
    // unsigned int delta_sold = yuan2fen(500);
    // unsigned int p = amm_get_delta_to_token_input_price(ex, delta_sold);
    // printf("price: %d\n", p);
    // ex->token_balance -= p;
    // for (int i=0; i<10; i++) {
    //     err = amm_delta_to_token_input_price(ex, yuan2fen(10), 1);
    //     if (err) {
    //         printf("amm_delta_to_token_input_price() = false\n");
    //         return 1;
    //     }
    // }
    unsigned int p = amm_get_delta_to_token_input_price(ex, yuan2fen(10));

    printf("amm_delta_to_token_input_price(1) = %d\n", p);
    printf("-- gettotal\n");
    printf("token_balance: %d\n", ex->token_balance);
    printf("total_supply: %d\n", ex->total_supply);
    printf("balance: %d\n", ex->balance);

    //amm_get_input_price(ex, );
    //amm_set_token_balance(ex, token_balance);
    //amm_get_input_price(ex, 1000, );
    ///ex->fn = &getfn;
    // amm_add_liquidity(ex, 1, 1, 1);
    //int n = tokena->balance;
    //printf("n=%s\n", tokena);
    //getfn();
    return 0;
}
