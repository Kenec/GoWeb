package viewModel

type whoWeAre struct {
	Description string
	Logo string
}

type whatWeDo struct{
	Text string
}

type interestUs struct {
	Phone int
	Address string
}

type Info struct {
	CompanyName string
	WhoWeAre whoWeAre
	WhatWeDo whatWeDo
	InterestUs interestUs
}

func NewInfo() Info {
	who := whoWeAre {
		Description: "This is a demo example of who we are as a company",
		Logo: "logo",
	}

	what := whatWeDo {
		Text: "What is do is learn Golang all day",
	}

	interest := interestUs {
		Phone: 0703,
		Address: "235 Ikorodu Rd, Ilupeju",
	}

	BitsCode := Info {
		CompanyName: "BitsCode - Nig",
		WhoWeAre: who,
		WhatWeDo: what,
		InterestUs: interest,
	}

	return BitsCode
}