package main

import (
	"fmt"
	"math/rand"
  "strings"
)

func main() {
var statu bool = true;
var name string;
var command string;
	for statu {
    fmt.Println("Panele Hoş geldin, İsim Alabilir miyim?");
    fmt.Scan(&name);
    fmt.Println("Hoş geldin", name);
    fmt.Println("Benden ne istersiniz efendim?");
    fmt.Println("- Hesap makinesi | 1");
    fmt.Println("- Oyunlar | 2");
    fmt.Scan(&command);
    switch command {
      case "1":
var chose string
 var num1 int
 var num2 int
 var result int
  
 fmt.Println("İşlem seçiniz");
 fmt.Println("- Toplama : +");
 fmt.Println("- Çıkarma : -");
 fmt.Println("- Çarpma : *");
 fmt.Println("- Bölme : /");
 fmt.Scan(&chose);

 switch(chose){
   case "+":
   fmt.Println("  1.Sayıyı giriniz");
   fmt.Scan(&num1);
   fmt.Println("  2.Sayıyı giriniz");
   fmt.Scan(&num2);
   result = num1 + num2;
   fmt.Println(result);
   break;
   case "-":
   fmt.Println("  1.Sayıyı giriniz");
   fmt.Scan(&num1);
   fmt.Println("  2.Sayıyı giriniz");
   fmt.Scan(&num2);
   result = num1 - num2;
   fmt.Println(result);
   break;
   case "*":
   fmt.Println("  1.Sayıyı giriniz");
   fmt.Scan(&num1);
   fmt.Println("  2.Sayıyı giriniz");
   fmt.Scan(&num2);
   result = num1 * num2;
   fmt.Println(result);
   break;
   case "/":
   fmt.Println("  1.Sayıyı giriniz");
   fmt.Scan(&num1);
   fmt.Println("  2.Sayıyı giriniz");
   fmt.Scan(&num2);
   result = num1 / num2;
   fmt.Println(result);
   break;
 }
  break;
  case "2":
   var gamechose string;
   fmt.Println("Oyun seçiniz");
   fmt.Println("- Taş kağıt makas | 1");
   fmt.Println("- Pokemon oyunu | 2");
   fmt.Println("- Maden oyunu | 3");
   fmt.Scan(&gamechose);
   switch gamechose {
     case "1":
var game bool = true;
  
  var chose string;
  var botchose int = rand.Intn(3);
  
  for game {
    fmt.Println("Taş : 1");
    fmt.Println("Kağıt : 2");
    fmt.Println("Makas : 3");
    
    fmt.Scan(&chose);
    botchose = rand.Intn(3);
    //Berabere durumları
    if chose == "1" && botchose == 1{
      fmt.Println("Sen : Taş");
      fmt.Println("Rakip : Taş");
      fmt.Println("Berabere !!");
    } else if chose == "2" && botchose == 2 {
      fmt.Println("Sen : Kağıt");
      fmt.Println("Rakip : Kağıt");
      fmt.Println("Berabere !!");      
    } else if chose == "3" && botchose == 3 {
      fmt.Println("Sen : Makas");
      fmt.Println("Rakip : Makas");
      fmt.Println("Berabere !!");      
    } else if chose == "1" && botchose == 3 { //Kazanma durumları
      fmt.Println("Sen : Taş");
      fmt.Println("Rakip : Makas");
      fmt.Println("Kazandın !!");      
    } else if chose == "2" && botchose == 1 {
      fmt.Println("Sen : Kağıt");
      fmt.Println("Rakip : Taş");
      fmt.Println("Kazandın !!");      
    } else if chose == "3" && botchose == 2 {
      fmt.Println("Sen : Makas");
      fmt.Println("Rakip : Kağıt");
      fmt.Println("Kazandın !!");      
    } else if chose == "1" && botchose == 2 { //Kaybetme durumları
      fmt.Println("Sen : Taş");
      fmt.Println("Rakip : Kağıt");
      fmt.Println("Kaybettin !!");      
    } else if chose == "2" && botchose == 3 {
      fmt.Println("Sen : Kağıt");
      fmt.Println("Rakip : Makas");
      fmt.Println("Kaybettin !!");      
    } else if chose == "3" && botchose == 1 {
      fmt.Println("Sen : Makas");
      fmt.Println("Rakip : Taş");
      fmt.Println("Kaybettin !!");       
    }
  }
     break;
     case "2":
	var name string
  var i bool = true
  var command string
  
  var animals = []string{}

  var animalList = []string{
  "Örümcek",
  "Ahtapot",
  "Kalamar",
  "Koyun",
  "İnek",
  "Salyangoz",
  "Aslan",
  "Kaplan",
  "Çita",
  "Jaguar",
  "Solucan",
  "Köstebek",
  "Fil",
  "Deve",
  "Arap",
  "Eşek",
  "At",
  "Unicorn",
  "Domuz",
  "Kelebek",
  "Baykuş",
  "Kartal",
  "Yengeç",
  "Sincap",
  "Kedi",
  "Papağan",
  "Güvercin",
  "Serçe",
  "Örümcek",
  "Ahtapot",
  "Kalamar",
  "Koyun",
  "İnek",
  "Salyangoz",
  "Aslan",
  "Kaplan",
  "Çita",
  "Jaguar",
  "Solucan",
  "Köstebek",
  "Fil",
  "Deve",
  "Arap",
  "Eşek",
  "At",
  "Unicorn",
  "Domuz",
  "Kelebek",
  "Baykuş",
  "Kartal",
  "Yengeç",
  "Sincap",
  "Kedi",
  "Papağan",
  "Güvercin",
  "Serçe",
  "Örümcek",
  "Ahtapot",
  "Kalamar",
  "Koyun",
  "İnek",
  "Salyangoz",
  "Aslan",
  "Kaplan",
  "Çita",
  "Jaguar",
  "Solucan",
  "Köstebek",
  "Fil",
  "Deve",
  "Arap",
  "Eşek",
  "At",
  "Unicorn",
  "Domuz",
  "Kelebek",
  "Baykuş",
  "Kartal",
  "Yengeç",
  "Sincap",
  "Kedi",
  "Papağan",
  "Güvercin",
  "Serçe",
  "Dinazor",
  }
 
  var para int = 0; 
  var rastpar int = 100;
  var olta bool = false;
  var ag bool = false;
  var upgpara int = 400;
  
  fmt.Println("-------------------");
  fmt.Println("    İsim nedir?    ");
  fmt.Println("-------------------");
  fmt.Scan(&name);
  for(i){
   fmt.Println("-------------------");
   fmt.Println("   Komut giriniz   ");
   fmt.Println("-------------------");
   fmt.Println("      KOMUTLAR     ");
   fmt.Println("-------------------");
   fmt.Println("");
   fmt.Println("avlan");   
   fmt.Println("hayvanlarım");
   fmt.Println("param");
   fmt.Println("market");
   fmt.Println("satin-al-esyaadi");  
   fmt.Println("güçlendir");
   fmt.Println("");
   fmt.Println("-------------------");
   fmt.Scan(&command);
   switch(strings.ToLower(command)){
     case "avlan":
     var kazanilanpara int = rand.Intn(rastpar);
     hunted := rand.Intn(len(animalList));
     if(olta == false){
     animals = append(animals, animalList[hunted]);
     para = para + kazanilanpara;
     } 
     if(olta == true) {
     hunted2 := rand.Intn(len(animalList));
     kazanilanpara = rand.Intn(rastpar)*2;
     animals = append(animals, animalList[hunted], animalList[hunted2]);
     fmt.Println("");
     fmt.Println("Tebrikler",animalList[hunted2],"Yakaladın!");
     fmt.Println("");   
     }
     if(ag == true){
      kazanilanpara = rand.Intn(rastpar);
     }
     
     animals = append(animals, animalList[hunted]);
     para = para + kazanilanpara;
     fmt.Println("");
     fmt.Println("Tebrikler",animalList[hunted],"Yakaladın!");
     fmt.Println("Kazanılan para :", kazanilanpara);
     fmt.Println("");
     
     break;
     case "hayvanlarım":
     fmt.Println("");
     fmt.Println(animals);
     fmt.Println("");
     break;
     case "param":
     fmt.Println("");
     fmt.Println(para,"TL");
     fmt.Println("");
     break;
     case "market":
     fmt.Println("");
     fmt.Println("- Ağ | Daha Çok para kazanmanızı sağlar [2x] | 200 TL");
     fmt.Println("- Olta | Daha çok hayvan yakalamanızı sağlar [2x] | 500 TL");
     fmt.Println("");
     break;
     case "satin-al-olta":
     if(para > 499){
     fmt.Println("");
     fmt.Println("Olta Satın alındı");
     olta = true;
     fmt.Println("");
     } else {
     fmt.Println("");
     fmt.Println("Yetersiz bakiye!");
     fmt.Println("");
     }
     break;
     case "satin-al-ag":
     if(para > 199){
     fmt.Println("");
     fmt.Println("Ağ Satın alındı");
     ag = true;
     rastpar = 200;
     fmt.Println("");
     } else {
     fmt.Println("");
     fmt.Println("Yetersiz bakiye!");
     fmt.Println("");
     }
     break;
     case "güçlendir":
     fmt.Println("");
     fmt.Println("para-arttır | Kazanılan parayı arttır | ",upgpara, "TL");
     fmt.Println("");
     break;
     case "para-arttır":
     if(para >= upgpara){
       upgpara = upgpara + rand.Intn(upgpara*4);
       rastpar = rastpar + rand.Intn(rastpar*2);
       fmt.Println("");
       fmt.Println("Para Arrttırma seviye atladı!");
       fmt.Println("Artık max",rastpar," Para kazanacaksın!")
       fmt.Println("");
     } else {
      fmt.Println("");
      fmt.Println("Yetersiz bakiye!");
      fmt.Println("");
     }
     break;
   }
  }
     break;
     case "3":
var game bool = true

	var madenler = []string{"Emrald", "Diamond", "Diamond", "Gold", "Gold", "Gold", "Iron", "Iron", "Iron", "Iron", "Bronz", "Silver", "Silver", "Silver", "Silver", "Silver", "Silver", "Silver", "Silver", "Silver", "Silver", "Silver", "Silver", "Coal", "Coal", "Coal", "Coal", "Coal", "Coal", "Coal", "Coal", "Coal", "Coal", "Coal", "Coal", "Coal", "Coal", "Coal", "Coal", "Coal", "Coal", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone"}

	var name string
	var command string
	var buy string
	var upgrade string

	var para int = 0
	var kazanc int = 100

	var kazilacaklar int = 1

	var mineupg int = 400
	var cashupg int = 200
	var minecountupg int = 200

	type picaxes struct {
		wooden  bool
		stone   bool
		iron    bool
		silver  bool
		gold    bool
		diamond bool
		emrald  bool
	}

	pickaxe := picaxes{
		wooden:  false,
		stone:   false,
		iron:    false,
		silver:  false,
		diamond: false,
		emrald:  false,
	}

	fmt.Println("What is your name gamer?");
	fmt.Scan(&name);
	for game {

		fmt.Println("_-_ Enter Command _-_");
		fmt.Scan(&command);
		switch command {
		case "mine":
			i := 1
			for i <= kazilacaklar {
				kazilan := madenler[rand.Intn(len(madenler))];
        switch(kazilan){
        case "Stone":
        kazancin := rand.Intn(kazanc);
				fmt.Println("WoW u find", kazilan, "!!");
				fmt.Println(kazancin, "money deposited");
				para = para + kazancin;
				i = i + 1;
        break;
        case "Coal":
        kazancin := rand.Intn(kazanc-40)*2;
				fmt.Println("WoW u find", kazilan, "!!");
				fmt.Println(kazancin, "money deposited");
				para = para + kazancin;
				i = i + 1;
        break;
        case "Iron":
        kazancin := rand.Intn(kazanc-40)*3;
				fmt.Println("WoW u find", kazilan, "!!");
				fmt.Println(kazancin, "money deposited");
				para = para + kazancin;
				i = i + 1;
        break;
        case "Silver":
        kazancin := rand.Intn(kazanc-40)*4;
				fmt.Println("WoW u find", kazilan, "!!");
				fmt.Println(kazancin, "money deposited");
				para = para + kazancin;
				i = i + 1;
        break;
        case "Gold":
        kazancin := rand.Intn(kazanc-40)*5;
				fmt.Println("WoW u find", kazilan, "!!");
				fmt.Println(kazancin, "money deposited");
				para = para + kazancin;
				i = i + 1;
        break;
        case "Diamond":
        kazancin := rand.Intn(kazanc-40)*6;
				fmt.Println("WoW u find", kazilan, "!!");
				fmt.Println(kazancin, "money deposited");
				para = para + kazancin;
				i = i + 1;
        break;
        case "Emrald":
        kazancin := rand.Intn(kazanc-40)*10;
				fmt.Println("WoW u find", kazilan, "!!");
				fmt.Println(kazancin, "money deposited");
				para = para + kazancin;
				i = i + 1;
        break;
        }
			}
			break
		case "market":
			fmt.Println("")
			fmt.Println("_-_  Pickaxes _-_")
			fmt.Println("1 >   Stone Pickaxes | 200 Coin");
			fmt.Println("2 >   Iron Pickaxe | 2000 Coin");
			fmt.Println("3 >   Silver Picaxe | 8000 Coin");
			fmt.Println("4 >   Gold Picaxe | 10000 Coin");
			fmt.Println("5 >   Diamond Picaxe | 17500 Coin");
			fmt.Println("6 >   Emrald Picaxe | 58000 Coin");
			fmt.Println("_-_ Upgrade List _-_");
			fmt.Println("7 >   Mining |", mineupg, "Coin");
			fmt.Println("8 >   Cash |", cashupg, "Coin");
			fmt.Println("9 >   Mining Count", minecountupg, "Coin");
			fmt.Println("")
			break
		case "upgrade":
			fmt.Println("_-_ Upgrade List _-_");
			fmt.Println("1 >   Mining |", mineupg, "Coin");
			fmt.Println("2 >   Cash |", cashupg, "Coin");
			fmt.Println("3 >   Mining Count", minecountupg, "Coin");
			fmt.Scan(&upgrade)
      switch(upgrade){
        case "1":
        if(para >= mineupg){
          para = para - mineupg;
          kazanc = kazanc + rand.Intn(kazanc);
          kazilacaklar = kazilacaklar + 1; 
          mineupg = mineupg + rand.Intn(mineupg*4);
          fmt.Println("Upgraded Cash, now you earn",kazanc,"per mine");
          fmt.Println("Mining Count upgraded, Now you mining",kazilacaklar,"Times!!");
        } else {
					fmt.Println("You have not money");          
        }
        break;
        case "2":
        if(para >= cashupg){
          para = para - cashupg;
          kazanc = kazanc + rand.Intn(kazanc);
          cashupg = cashupg + rand.Intn(cashupg);
          fmt.Println("Upgraded Cash, now you earn",kazanc,"per mine");
        } else {
 					fmt.Println("You have not money");         
        }
        break;
        case "3":
        if(para >= minecountupg){
          para = para - minecountupg;
          kazilacaklar = kazilacaklar + 1;
          minecountupg = minecountupg + rand.Intn(minecountupg);
          fmt.Println("Mining Count upgraded, Now you mining",kazilacaklar,"Times!!");
        } else {
					fmt.Println("You have not money");          
        }
        break;
      }
			break
		case "buy":
			fmt.Println("Select a Item");
			fmt.Println("If you don't know what to buy try the market command");
			fmt.Println("1 >   Stone Pickaxes | 200 Coin");
			fmt.Println("2 >   Iron Pickaxe | 2000 Coin");
			fmt.Println("3 >   Silver Picaxe | 8000 Coin");
			fmt.Println("4 >   Gold Picaxe | 10000 Coin");
			fmt.Println("5 >   Diamond Picaxe | 17500 Coin");
			fmt.Println("6 >   Emrald Picaxe | 58000 Coin");
			fmt.Scan(&buy)
			switch buy {
			case "1":
				if para >= 200 && pickaxe.stone == false {
					pickaxe.stone = true;
					para = para - 200;
          kazilacaklar = kazilacaklar + 1;
					fmt.Println("You Got a Stone Pickaxe");
				} else {
					fmt.Println("You have not money");
				}
				break
			case "2":
				if para >= 2000 && pickaxe.iron == false {
					pickaxe.iron = true;
          para = para - 2000;
          kazilacaklar = kazilacaklar + 2;
					fmt.Println("You Got a Iron Pickaxe");
				} else {
					fmt.Println("You have not money");
				}
				break
			case "3":
				if para >= 8000 && pickaxe.silver == false {
					pickaxe.silver = true;
          para = para - 8000;
          kazilacaklar = kazilacaklar + 3;
					fmt.Println("You Got a Sılver Pickaxe");
				} else {
					fmt.Println("You have not money");
				}
				break
			case "4":
				if para >= 10000 && pickaxe.gold == false {
          pickaxe.gold = true;
          para = para - 10000;
          kazilacaklar = kazilacaklar + 4;
					fmt.Println("You Got a Gold Pickaxe");
				} else {
					fmt.Println("You have not money");
				}
				break
			case "5":
				if para >= 17500 && pickaxe.diamond == false {
          pickaxe.diamond = true;
          para = para - 175000;
          kazilacaklar = kazilacaklar + 6;
					fmt.Println("You Got a Diamond Pickaxe");
				} else {
					fmt.Println("You have not money");
				}
				break
			case "6":
				if para >= 58000 && pickaxe.emrald == false {
          pickaxe.emrald = true;
          para = para - 58000;
          kazilacaklar = kazilacaklar + 10;
					fmt.Println("You Got a Emrald Pickaxe");
				} else {
					fmt.Println("You have not money");
				}
				break
			}
			break
		case "cash":
			fmt.Println("You have a", para, "Coin");
			break
		}
	}
     break;
   }
  break;
  }
 }
}
