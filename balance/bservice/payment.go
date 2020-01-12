package bservice

/*

import (
    "errors"
    "github.com/sample-restaurant-rest-api/entity"
    "time"
)


import(
"time"
"errors"
"../../entity"
)
func CheckSeller (S entity.User,b BalanceService) (entity.Balance, errors) {
ba,err:=b.Balance(S.id)
if err != nil{
return ba,err
}
return ba,nil
}
func CheckBuyer(u entity.User,pri double,b BalanceService)  (entity.Balance,errors){
ba,err :=b.Balance(u.id)
if err != nil{
return ba,err
}
if (ba.balance - pri - 0.5*ba.balance) < 0{
err:=errors.New("low account")
return ba,err
}
return ba,nil

}
func Payment (p entity.Product,buyer entity.User,b BalanceService) errors{
baB,err1:=CheckBuyer(buyer,p.price,b)
bas,err2:=CheckSeller(p.USER,b)
if (err1!=nil) || (err2!=nil){
return errors.New("Cannot make Payment")
}
bas.balance +=p.price - 0.5*bas.balance
baB.balance -=p.price - 0.5*baB.balance
err1:=b.Update(baB)
err2:=b.Update(bas)
                        i f(err1!=nil)||(err2!=nil){
return errors.New("Could not make Payment")
}
//storePaymentdetail(u.id,p,b)
return nil
}




//or


func paymentseller (p entity.Product,buyer entity.User,b BalanceService) errors{
bas,err2:=CheckSeller(p.USER,b)
baB,err1:=CheckBuyer(buyer,p.price,b)
if (err1!=nil) || (err2!=nil){
return errors.New("Cannot make Payment")
}

//run as go routine
bas.Balance +=p.price - 0.5*bas.Balance
sleep(60*time.Second)
err2:=b.Update(bas)
}

*/
