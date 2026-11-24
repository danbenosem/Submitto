package repositories_test


import ("testing"


		"github.com/stretchr/testify/assert"
		"github.com/stretchr/testify/require"
			"submitto/internals/database"
			"submitto/internals/models"
)


func  testThatFormCanBeCreated(t *testing.T) {

   db, err := database.Connect()
   require.NoError(t, err)

   err =db.AutoMigrate(&models.Form{})

   require.NoError(t, err)

   form:= models.Form{
		UserId: 1,
		Title: "Test Form",
		Description: "This is a test form",
   }

}