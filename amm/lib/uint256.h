#ifndef UINT256_H_
#define UINT256_H_

#include <stdbool.h>
#include <stdint.h>

#define _UINT256_LEN 8

#define UINT256_ZERO uint256_new()
#define UINT256_ONE uint256_set_uint8(UINT256_ZERO, 1)
#define UINT256_TWO uint256_set_uint8(UINT256_ZERO, 2)
#define UINT256_TEN uint256_set_uint8(UINT256_ZERO, 10)

#define UINT256_U8(n) uint256_set_uint8(UINT256_ZERO, n)
#define UINT256_U32(n) uint256_set_uint32(UINT256_ZERO, n)
#define UINT256_U64(n) uint256_set_uint64(UINT256_ZERO, n)
#define UINT256(n) UINT256_U32(n)

struct uint256 {
    uint64_t ns[_UINT256_LEN];
};

typedef struct uint256 uint256_t;


uint256_t *uint256_new();
uint256_t *uint256_set_uint8(uint256_t *, uint8_t);
uint256_t *uint256_set_uint32(uint256_t *, uint32_t);
uint256_t *uint256_set_uint64(uint256_t *, uint64_t);

bool uint256_is_zero(uint256_t *);
int uint256_cmp(uint256_t *, uint256_t *);
bool uint256_eq(uint256_t *, uint256_t *);
bool uint256_lt(uint256_t *, uint256_t *);
bool uint256_lte(uint256_t *, uint256_t *);
bool uint256_gt(uint256_t *, uint256_t *);
bool uint256_gte(uint256_t *, uint256_t *);

// addition
uint256_t *uint256_add(uint256_t *, uint256_t *);
// subtraction
uint256_t *uint256_sub(uint256_t *, uint256_t *);
// multiply
uint256_t *uint256_mul(uint256_t *, uint256_t *);
// divide
uint256_t *uint256_div(uint256_t *, uint256_t *);

void uint256_to_hex(uint256_t *, char *, uint32_t);
void uint256_to_string(uint256_t *, char *, uint32_t);

#endif
