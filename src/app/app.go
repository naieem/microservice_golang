package app

import (
  "github.com/gin-gonic/gin"
)

func GetInstructions(c *gin.Context){
  selDB, err := dbmap.Query("SELECT title,description FROM articles")
  if err != nil {
    c.JSON(404, gin.H{"error": err.Error()})
  }else{
  // Call the struct to be rendered on template
  n := Instruction{}

  // Create a slice to store all data from struct
  res := []Instruction{}

  // Read all rows from database
    for selDB.Next() {
      // Must create this variables to store temporary query
      // var id int
      var title, description string

      // Scan each row storing values from the variables above and check for errors
      err = selDB.Scan(&title, &description)
      if err != nil {
        panic(err.Error())
      }

      // Get the Scan into the Struct
      // n.Id = id
      n.Title = title
      n.Description = description

      // Join each row on struct inside the Slice
      res = append(res, n)

    }
    c.JSON(200, res)
  }
}

func UpdateInstruction(c *gin.Context) {
  c.JSON(200, gin.H{"ok": "PUT api/v1/instructions/1"})

}

func DeleteInstruction(c *gin.Context) {
  c.JSON(200, gin.H{"ok": "DELETE api/v1/instructions/1"})
}

// func GetInstruction(c *gin.Context) {
//   id := c.Params.ByName("id")
//   var instruction Instruction
  
//   err := dbmap.SelectOne(&instruction, "SELECT * FROM instruction WHERE id=?", id)
//   if err == nil {
//     // instructionID, _ := strconv.ParseInt(id, 0, 64)
  
//     content := &Instruction{
//       Title: instruction.Title,
//       Description: instruction.Description,
//     }
 
//     c.JSON(200, content)
//   } else {
//     c.JSON(404, gin.H{"error": "instruction not found"})
//   }
//   // curl -i http://localhost:8080/api/v1/Instructions/1
// }
