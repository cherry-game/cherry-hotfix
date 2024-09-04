module github.com/zc-deng/cherry-hotfix

go 1.21

toolchain go1.22.6

require (
	github.com/agiledragon/gomonkey/v2 v2.10.1
	github.com/cherry-game/cherry-hotfix v0.0.0-20240901042442-bd4b2a166736
	github.com/traefik/yaegi v0.15.1
)

replace github.com/cherry-game/cherry-hotfix => github.com/zc-deng/cherry-hotfix v0.0.0-20240904025944-98130738a0b1
replace github.com/zc-deng/cherry-hotfix => github.com/zc-deng/hotfix v0.0.0-20240904025944-98130738a0b1
