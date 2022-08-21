package main

func main() {
	(&Bot{}).Start(getenv("TELEGRAM_TOKEN"), getenv("COINRANKING_TOKEN"))
}
