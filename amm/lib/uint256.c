#include "uint256.h"
#include <stdio.h>
#include <stdlib.h>

uint256_t *uint256_new() {
    uint256_t *np = malloc(sizeof(uint256_t));
    return np;
}

uint256_t *uint256_set_uint8(uint256_t *u, uint8_t n) {
    u->ns[0] = (uint64_t) n;
    return u;
}

uint256_t *uint256_set_uint32(uint256_t *u, uint32_t n) {
    u->ns[0] = (uint64_t) n;
    return u;
}

uint256_t *uint256_set_uint64(uint256_t *u, uint64_t n) {
    u->ns[1] = (n & 0xffffffff00000000) >> 32;
    u->ns[0] = n & 0xffffffff;
    return u;
}

int uint256_cmp(uint256_t *a, uint256_t *b){
    
    return 0;
}

uint256_t *uint256_mul(uint256_t *a, uint256_t *b){
    uint64_t c = 0;
    for (int i=0; i<_UINT256_LEN; i++){
        uint64_t ai = a->ns[i];
        uint64_t bi = b->ns[i];
        uint64_t m = c * ai * bi;
        a->ns[i] = m & 0xffffffff;
        c = m >> 32;
    }
    return a;
}
uint256_t *uint256_add(uint256_t *a, uint256_t *b){
    uint64_t c = 0;
    for (int i=0; i<_UINT256_LEN; i++){
        uint64_t ai = a->ns[i];
        uint64_t bi = b->ns[i];
        uint64_t m = c + ai + bi;
        a->ns[i] = m & 0xffffffff;
        c = m >> 32;
    }
    return a;
}
void uint256_to_hex(uint256_t *u, char *str, uint32_t len) {
    const char cs[16] = "0123456789abcdef";
    for (int i=0; i<_UINT256_LEN; i++) {
        uint32_t n = u->ns[i];
        char hex[9];
        sprintf(hex, "%08x", n);
        printf("hex=%s\n", hex);
    }
}
void uint256_to_string(uint256_t *u, char *str, uint32_t len) {
    for (int i=0; i<_UINT256_LEN; i++){
        uint64_t n = u->ns[i];
        printf("n=%llu\n", n);
    }
}
