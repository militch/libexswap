CREATE TABLE `market_days_history` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT '交易对ID', 
    `pool_name` varchar(150) NOT NULL COMMENT '交易对名称', 
    `token_main` varchar(150) NOT NULL COMMENT '主要的通证符号', 
    `token_second` varchar(150) NOT NULL COMMENT '次要的通证符号',
    `main_token_decimals` tinyint(1) NOT NULL DEFAULT '18' COMMENT '通证精度(主)',
    `second_token_decimals` tinyint(1) NOT NULL DEFAULT '18' COMMENT '通证精度(次)',
    `amount` varchar(150) NOT NULL DEFAULT '0' COMMENT '以主要通证计量的交易量',
    `volume` varchar(150) NOT NULL DEFAULT '0' COMMENT '以次要通证计量的交易量',
    `trade_count` int NOT NULL DEFAULT '0' COMMENT '交易次数',
    `avg_liquidity` varchar(150) NOT NULL DEFAULT '0' COMMENT '平均流动性',
    `open_price` varchar(150) NOT NULL DEFAULT '0' COMMENT '开盘价',
    `avg_price` varchar(150) NOT NULL DEFAULT '0' COMMENT '平均价格',
    `close_price` varchar(150) NOT NULL DEFAULT '0' COMMENT '收盘价',
    `lower_price` varchar(150) NOT NULL DEFAULT '0' COMMENT '最低价',
    `highest_price` varchar(150) NOT NULL DEFAULT '0' COMMENT '最高价',
    `create_time` datetime NOT NULL DEFAULT current_timestamp COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT current_timestamp COMMENT '更新时间',
    KEY (`id`)
) COMMENT='交易市场按日统计数据';

CREATE TABLE `market_monthly_history` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT '交易对ID', 
    `pool_name` varchar(150) NOT NULL COMMENT '交易池名称', 
    `token_main` varchar(150) NOT NULL COMMENT '主要的通证符号', 
    `token_second` varchar(150) NOT NULL COMMENT '次要的通证符号',
    `main_token_decimals` tinyint(1) NOT NULL DEFAULT '18' COMMENT '通证精度(主)',
    `second_token_decimals` tinyint(1) NOT NULL DEFAULT '18' COMMENT '通证精度(次)',
    `amount` varchar(150) NOT NULL DEFAULT '0' COMMENT '以主要通证计量的交易量',
    `volume` varchar(150) NOT NULL DEFAULT '0' COMMENT '以次要通证计量的交易量',
    `trade_count` int NOT NULL DEFAULT '0' COMMENT '交易次数',
    `avg_liquidity` varchar(150) NOT NULL DEFAULT '0' COMMENT '平均流动性',
    `open_price` varchar(150) NOT NULL DEFAULT '0' COMMENT '开盘价',
    `avg_price` varchar(150) NOT NULL DEFAULT '0' COMMENT '平均价格',
    `close_price` varchar(150) NOT NULL DEFAULT '0' COMMENT '收盘价',
    `lower_price` varchar(150) NOT NULL DEFAULT '0' COMMENT '最低价',
    `highest_price` varchar(150) NOT NULL DEFAULT '0' COMMENT '最高价',
    `create_time` datetime NOT NULL DEFAULT current_timestamp COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT current_timestamp COMMENT '更新时间',
    KEY (`id`)
) COMMENT='交易市场按月统计数据';

CREATE TABLE `account_token_assets` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT '资产id',
    `user_id` varchar(150) NOT NULL COMMENT '用户ID',
    `token_symbol` varchar(150) NOT NULL COMMENT '通证符号',
    `token_decimals` tinyint(1) NOT NULL DEFAULT '18' COMMENT '通证精度',
    `token_balance` varchar(150) NOT NULL COMMENT '通证余额',
    `create_time` datetime NOT NULL DEFAULT current_timestamp COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT current_timestamp COMMENT '更新时间',
    KEY (`id`)
) COMMENT='账户通证资产';

CREATE TABLE `account_assets_log` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT '账户资产流水记录ID', 
    `log_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '记录类型(0=充值记录, 1=转移记录, 2=交易记录, 3=流动池交易)',
    `log_id` int NULL DEFAULT NULL COMMENT '业务记录ID',
    `token_symbol` varchar(150) NOT NULL COMMENT '通证符号',
    `token_decimals` tinyint(1) NOT NULL DEFAULT '18' COMMENT '通证精度',
    `from_token_balance` varchar(150) NOT NULL COMMENT '变更前资产余额',
    `to_token_balance` varchar(150) NOT NULL COMMENT '变更后资产余额',
    `token_amount` varchar(150) NOT NULL COMMENT '通证金额',
    `user_id` varchar(150) NOT NULL COMMENT '用户ID',
    `create_time` datetime NOT NULL DEFAULT current_timestamp COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT current_timestamp COMMENT '更新时间',
    KEY (`id`)
) COMMENT='账户资产流水表';

CREATE TABLE `exchange_token` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT '通证ID', 
    `name` varchar(150) NOT NULL COMMENT '通证名称', 
    `symbol` varchar(150) NOT NULL COMMENT '通证符号',
    `icon_url` varchar(150) NULL DEFAULT '' COMMENT '通证图标',
    `decimals` tinyint(1) NOT NULL DEFAULT '18' COMMENT '通证精度',
    `biz_decimals` tinyint(1) NOT NULL DEFAULT '6' COMMENT '业务展示精度',
    `description` varchar(150) NULL DEFAULT '' COMMENT '描述',
    `is_disabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否禁用(0=FALSE, 1=TRUE)',
    `weight` int NOT NULL DEFAULT '0' COMMENT '排序权重',
    `create_time` datetime NOT NULL DEFAULT current_timestamp COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT current_timestamp COMMENT '更新时间',
    KEY (`id`)
) COMMENT='通证配置';

CREATE TABLE `exchange_token_pair` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT '交易对ID', 
    `name` varchar(150) NOT NULL COMMENT '交易对名称', 
    `token_main` varchar(150) NOT NULL COMMENT '主要的通证符号', 
    `token_second` varchar(150) NOT NULL COMMENT '次要的通证符号',
    `description` varchar(150) NULL DEFAULT '' COMMENT '描述',
    `is_disabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否禁用(0=FALSE, 1=TRUE)',
    `weight` int NOT NULL DEFAULT '0' COMMENT '排序权重',
    `create_time` datetime NOT NULL DEFAULT current_timestamp COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT current_timestamp COMMENT '更新时间',
    KEY (`id`)
) COMMENT='交易对配置';

CREATE TABLE `exchange_pool` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT '交易池ID', 
    `name` varchar(150) NOT NULL COMMENT '交易池名称', 
    `token_main` varchar(150) NOT NULL COMMENT '主要的通证符号', 
    `token_second` varchar(150) NOT NULL COMMENT '次要的通证符号',
    `main_balance` varchar(150) NOT NULL DEFAULT '0' COMMENT '交易池余额(主)',
    `second_balance` varchar(150) NOT NULL DEFAULT '0' COMMENT '交易池余额(次)',
    `total_supply` varchar(150) NOT NULL DEFAULT '0' COMMENT '总供给量(流动性)',
    `create_time` datetime NOT NULL DEFAULT current_timestamp COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT current_timestamp COMMENT '更新时间',
    KEY (`id`)
) COMMENT='交易池';

CREATE TABLE `exchange_pool_assets` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT '资产记录ID', 
    `pool_id` int NOT NULL COMMENT '交易池ID',
    `pool_name` varchar(150) NOT NULL COMMENT '交易池名称', 
    `token_main` varchar(150) NOT NULL COMMENT '主要的通证符号', 
    `token_second` varchar(150) NOT NULL COMMENT '次要的通证符号',
    `user_id` varchar(150) NOT NULL COMMENT '用户ID',
    `balance` varchar(150) NOT NULL DEFAULT '0' COMMENT '账户额度',
    `create_time` datetime NOT NULL DEFAULT current_timestamp COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT current_timestamp COMMENT '更新时间',
    KEY (`id`)
) COMMENT='交易池资产表';

CREATE TABLE `exchange_pool_assets_log` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT '交易池资产流水记录ID', 
    `pool_id` int NOT NULL COMMENT '交易池ID',
    `pool_name` varchar(150) NOT NULL COMMENT '交易池名称', 
    `token_main` varchar(150) NOT NULL COMMENT '主要的通证符号', 
    `token_second` varchar(150) NOT NULL COMMENT '次要的通证符号',
    `log_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '记录类型(0=增加流动资产, 1=移除流动资产)',
    `main_token_decimals` tinyint(1) NOT NULL DEFAULT '18' COMMENT '通证精度(主)',
    `second_token_decimals` tinyint(1) NOT NULL DEFAULT '18' COMMENT '通证精度(主)',
    `from_token_balance` varchar(150) NOT NULL COMMENT '变更前资产余额',
    `to_token_balance` varchar(150) NOT NULL COMMENT '变更后资产余额',
    `from_pool_balance` varchar(150) NOT NULL COMMENT '变更前交易池资产余额',
    `to_pool_balance` varchar(150) NOT NULL COMMENT '变更后交易池池资产余额',
    `from_pool_total_supply` varchar(150) NOT NULL COMMENT '变更前交易池流动性',
    `to_pool_total_supply` varchar(150) NOT NULL COMMENT '变更后交易池流动性',
    `token_amount` varchar(150) NOT NULL COMMENT '通证金额',
    `lp_amount` varchar(150) NOT NULL COMMENT '交易池金额',
    `user_id` varchar(150) NOT NULL COMMENT '用户ID',
    `create_time` datetime NOT NULL DEFAULT current_timestamp COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT current_timestamp COMMENT '更新时间',
    KEY (`id`)
) COMMENT='交易池资产流水表';


CREATE TABLE `exchange_trade_log` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT '交易记录ID', 
    `log_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '记录类型(0=买入记录, 1=卖出记录)',
    `pool_id` int NOT NULL COMMENT '交易池ID',
    `pool_name` varchar(150) NOT NULL COMMENT '交易池名称', 
    `token_main` varchar(150) NOT NULL COMMENT '主要的通证符号', 
    `token_second` varchar(150) NOT NULL COMMENT '次要的通证符号',
    `price` varchar(150) NOT NULL COMMENT '成交价格',
    `token_amount` varchar(150) NOT NULL COMMENT '通证数量',
    `user_id` varchar(150) NOT NULL COMMENT '用户ID',
    `from_token_balance` varchar(150) NOT NULL COMMENT '变更前资产余额',
    `to_token_balance` varchar(150) NOT NULL COMMENT '变更后资产余额',
    `from_pool_total_supply` varchar(150) NOT NULL COMMENT '变更前交易池流动性',
    `to_pool_total_supply` varchar(150) NOT NULL COMMENT '变更后交易池流动性',
    `create_time` datetime NOT NULL DEFAULT current_timestamp COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT current_timestamp COMMENT '更新时间',
    KEY (`id`)
) COMMENT='交易记录表';
