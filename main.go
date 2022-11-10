package main

import (
	"fmt"
  "math/rand"
  "strings"
  "time"
)

func main() {
	var name string
  var chose string
  var hchose string
  var city1 string
  var city2 string
  
  var number1 int
  var number2 int
  var tax int
  var cash int
  var diff int
  var result int

  var info = []string{
    "Balık krakerleri hepimiz biliriz. İşte bunlar tek tip değiller: Bazı balık krakerler sağa bakarken bazıları sola bakarlar.",
    "Bu bilgiyi sosyal medyada dolanırken görmüş olmanız çok olası: “PACIFIC OCEAN” kelimesindeki her “C” birbirinden farklı telaffuz ediliyor.",
    "Kutup ayısının karaciğerinde o kadar yüksek oranda A vitamini vardır ki, eğer yerseniz yüksek dozdan ölürsünüz.",
    "Futbolcular ortalama bir maç sırasında yaklaşık 11 km koşarlar.",
    "TEVFİK FİKRET adındaki ilk hece ile soyadındaki ilk hece birleştirildiğinde adı olan; TEV-FİK, adındaki son hece ile soyadındaki son hece birleştirildiğinde de soyadı olan FİKRET meydana geliyor.",
    "Herkesin bayılarak izlediği Aşk-ı Memnu dizisindeki Nihal karakteri, dizi boyunca tam 37 kere bayılmıştır.",
    "Eğer bir su aygırı üzülürse terleri kırmızı renkli olur.",
    "Bu bilgi biraz endişe verici: Her gün yeni doğan 12 bebek yanlış aileye teslim ediliyor.",
    "Brezilya’da dünyanın en geniş yolu bulunmaktadır ve bu yolda 160 tane araba yan yana gidebilirler.",
    "Hamam böceklerinin kafaları kopsa bile yaklaşık birkaç hafta daha yaşamaya devam edebilirler.",
    "Renklendiriciler kullanılmasaydı, kolanın orijinal rengi yeşildi.",
    "Suudi Arabistan’daki kadınların eşleri onlardan kahve isterse ve kadın kocasına kahve yapmazsa bu bir boşanma sebebidir.",
    "İguanaların su altında kalabilme süreleri yaklaşık 28 dakikadır.",
    "İnsanlar doğduklarında 300 kemikleri varken yetişkin bir birey olduklarında bu sayı 206’ya düşer. Hatta bu kemiklerin yaklaşık 4’te 1’i ayaklardadır.",
    "İnekler, müziği çok severler. Hatta inekler müzik dinlerse daha fazla süt verirler.",
    "Amerika’da yaşayan köpeklerin sayısı 52 milyondan fazladır.",
    "Dünyadaki tüm tavukların sayısı ise insan sayısından fazladır.",
  }

 

  var i bool = true;

  fmt.Println("Merhabalar isim rica edebilir miyim? -",time.Now().Hour(),":", time.Now().Minute());
  fmt.Scan(&name);
  for(i){
  if(time.Now().Hour() < 22){
  fmt.Println("Hoş geldin",name, " sana nasıl yardımcı olabilirim?");
  fmt.Println("- Hesap makinesi : 1");
  fmt.Println("- Vergi Hesaplama : 2");
  fmt.Println("- Zar at : 3");
  fmt.Println("- Gereksiz Bilgi : 4");
  fmt.Println("- Esenler otogarı : 5");
  fmt.Println("- Saat kaç? : 6");
    
  fmt.Scan(&chose);
  switch(chose){
    case "1":
    fmt.Println("İşlem seçiniz");
    fmt.Println("- Toplama : +");
    fmt.Println("- Çıkarma : -");
    fmt.Println("- Bölme : /");
    fmt.Println("- Çarpma : *");
    fmt.Scan(&hchose);
     switch(hchose){
      case "+":
      fmt.Println("- 1.Sayıyı giriniz");
      fmt.Scan(&number1);
      fmt.Println("- 2.Sayıyı giriniz");
      fmt.Scan(&number2);
      result = number1 + number2;
      fmt.Println("Sonuç :",result);
      break;
      case "-":
      fmt.Println("- 1.Sayıyı giriniz");
      fmt.Scan(&number1);
      fmt.Println("- 2.Sayıyı giriniz");
      fmt.Scan(&number2);
      result = number1 - number2;
      fmt.Println("Sonuç :",result);
      break;
      case "/":
      fmt.Println("- 1.Sayıyı giriniz");
      fmt.Scan(&number1);
      fmt.Println("- 2.Sayıyı giriniz");
      fmt.Scan(&number2);
      result = number1 / number2;
      fmt.Println("Sonuç :",result);
      break;
      case "*":
      fmt.Println("- 1.Sayıyı giriniz");
      fmt.Scan(&number1);
      fmt.Println("- 2.Sayıyı giriniz");
      fmt.Scan(&number2);
      result = number1 * number2;
      fmt.Println("Sonuç :",result);
      break;
     }
    break;
    case "2":
     fmt.Println("Gelir giriniz");
     fmt.Scan(&cash);
     if(cash <= 12600){
       tax = cash*15/100;
       fmt.Println("Ödenecek vergi :",tax);
     } 
     if(cash > 12600 && cash <= 30000){
       diff = cash - 12600;
       tax = 1890 + diff * 20 / 100;
       fmt.Println("Ödenecek vergi :", tax);
     }
     if(cash > 30000 && cash <= 69000){
       diff = cash - 30000;
       tax = 5370 + diff * 27 / 100;
       fmt.Println("Ödenecek vergi :", tax);
     }
    if(cash > 69000 && cash <= 110000){
       diff = cash - 69000;
       tax = 15900 + diff * 35 / 100;
       fmt.Println("Ödenecek vergi :", tax); 
    }
    if(cash > 110000){
       diff = cash - 110000;
       tax = 26970 + diff * 35 / 100;
       fmt.Println("Ödenecek vergi :", tax); 
    }
    break;
    case "3":
     fmt.Println("Zar Atıldı, gelen sayı : ",rand.Intn(6))
    break;
    case "4":
      fmt.Println("Günün gereksiz bilgisi");
      fmt.Println("----------------------");
      fmt.Println(info[rand.Intn(len(info))]);
    break;
    case "5":
      fmt.Println("Hangi şehirdesiniz?");
      fmt.Scan(&city1);
      fmt.Println("Hangi Şehre gideceksiniz?");
      fmt.Scan(&city2);
      if(strings.ToLower(city1) == "konya" || strings.ToLower(city2) == "konya"){
        fmt.Println("Konya AsdAsdAsdAsdAsd:DSd")
      } else {
        fmt.Println(city1,"'den",city2,"ye bilet yok maalesef");
      }
    break;
    case "6":
     fmt.Println(time.Now().Hour(),":", time.Now().Minute());
    break;
  }
} else {
  fmt.Println("Bu saatte ne istiyorsun aq git yat");
   i = false;
  }
 }
} 
