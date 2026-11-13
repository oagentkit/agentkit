# oagentkit/agentkit —— 退役仓（v1.9.0 终版）：本地门禁与 CI(retired-contract)同口径
# 仓里没有代码，门禁守的是「墓碑形状」：零 .go、零依赖、README 刷正。
.PHONY: check retired-shape readme-fresh help

## check: 跑全部退役契约检查（与 CI retired-contract 同口径）
check: retired-shape readme-fresh

## retired-shape: 零 .go 文件 + go.mod 只许 module 声明
retired-shape:
	@bad=$$(find . -name '*.go' -not -path './.git/*'); if [ -n "$$bad" ]; then \
		echo "退役仓出现 .go 文件："; echo "$$bad"; exit 1; fi
	@if grep -qE '^[[:space:]]*(require|exclude|replace)[ (]' go.mod; then \
		echo "go.mod 出现 require/exclude/replace——墓碑长出了依赖"; exit 1; fi
	@grep -q '^module github.com/oagentkit/agentkit$$' go.mod || \
		{ echo "go.mod 缺 module 声明"; exit 1; }
	@echo "OK：零 .go 文件，go.mod = module 声明 + 零依赖"

## readme-fresh: README 状态行不落后最新 tag（沿用 #192 刷正纪律）
readme-fresh:
	@latest=$$(git describe --tags --abbrev=0); \
	readme=$$(grep -oE '当前 v[0-9]+\.[0-9]+\.[0-9]+' README.md | head -n1 | sed 's/^当前 //'); \
	if [ -z "$$readme" ]; then echo "README 状态行找不到「当前 vX.Y.Z」"; exit 1; fi; \
	if [ "$$(printf '%s\n%s\n' "$$latest" "$$readme" | sort -V | head -n1)" != "$$latest" ]; then \
		echo "README 状态行 $$readme 落后最新 tag $$latest"; exit 1; fi
	@echo "OK：README 状态行 ≥ 最新 tag"

## help: 显示本文件的目标
help:
	@grep -E '^## ' $(MAKEFILE_LIST) | sed 's/## //'
