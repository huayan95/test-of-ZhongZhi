package main

import (
    db "project1/database" 
router "project1/routers"
)

func main() {
   defer db.SqlDB.Close()
   router:=router.InitRouter()
   router.Static("/static", "./static")
   router.Run(":8000")
}
