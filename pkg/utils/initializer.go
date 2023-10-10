package utils

import (
	"github.com/ARUNK2121/procast/pkg/config"
	"github.com/ARUNK2121/procast/pkg/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CheckAndCreateAdmin(db *gorm.DB, cfg config.Config) {
	//check if any rows available at the admin table
	//if there is no existing admin create a super admin
	var count int64
	db.Model(&domain.Admin{}).Count(&count)
	if count == 0 {
		password := cfg.ADMIN_PASSWORD
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return
		}

		admin := domain.Admin{
			ID:        1,
			Name:      "procast",
			Email:     cfg.ADMIN_EMAIL,
			Password:  string(hashedPassword),
			Previlege: "super_admin",
		}
		db.Create(&admin)
	}
}

var Description string = `Introduction
Welcome to Procast, where skills meets opportunity! Procast is a game-changing aggregator business model operating in the dynamic realm of C2C (Consumer-to-Consumer) interactions. Our mission is to empower individuals with remarkable skills, connecting them to a vast marketplace of opportunities they might never have discovered otherwise.

How Procast Works

LISTING OPENINGS :
	 Are you in need of skilled services? Procast makes it easy to find the talent you require. Simply list your job openings on our platform, whether you need a web designer, a content writer, a graphic artist, or any other service.

RECIEVE BIDS :
	 Once your job opening is listed, skilled service providers from diverse backgrounds and expertise levels have the opportunity to bid on your project. These bids include details about the provider's experience, pricing, and proposed timeline.

CHOOSE THE BEST PROVIDER :
	 Review all the bids you receive and select the provider that best aligns with your requirements, budget, and quality expectations. Our platform makes it easy to compare and contrast bids, ensuring that you make an informed decision.

Why Procast?

ACCESS TO A VAST TALENT POOL :
	 Procast opens doors to a wealth of talent you might not have encountered through traditional channels. Our platform is home to a diverse array of skilled professionals, from freelancers and solopreneurs to experts with years of experience.

STREAMLINED PROCESS :
 	Say goodbye to the hassles of searching for contacts and the uncertainty of finding the right person for the job. Procast simplifies the process by bringing opportunities directly to your fingertips.

TRANSPARENCY AND QUALITY ASSURANCE : 
	Procast fosters transparency by providing comprehensive information about service providers, including reviews and ratings from previous clients. This way, you can make decisions based on past performance and the quality of work.

BUDGET-FRIENDLY OPTIONS :
	Procast accommodates a range of budgets, ensuring that you find the right provider without breaking the bank.`
