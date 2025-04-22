package main
import "fmt"
const NMAX int = 1024
type satuanWaktu struct{
	menit, detik int
}
type arrWaktu[NMAX]satuanWaktu
func main(){
	var waktuPembalap arrWaktu
	var nPembalap, i, waktuMinimal, menitMulai, detikMulai int
	fmt.Scan(&nPembalap)
	for i = 0; i < nPembalap; i++{
		fmt.Scan(&menitMulai, &detikMulai)
		waktuPembalap[i] = satuanWaktu{menitMulai, detikMulai}
	}
	fmt.Scan(&menitMulai, &detikMulai)
	waktuMinimal = waktuTercepat(waktuPembalap, nPembalap, menitMulai, detikMulai)
	fmt.Print(waktuMinimal)
}
func waktuTercepat(waktuPembalap arrWaktu, nPembalap, menitMulai, detikMulai int)int{
	var i, waktuMulai, p, mindetik, totalwaktu, banding int
	waktuMulai = (60*menitMulai)+detikMulai
	p = waktuPembalap[0].menit*60 + waktuPembalap[0].detik
	mindetik = waktuMulai-p
	for i = 0; i < nPembalap; i++{
		totalwaktu = waktuPembalap[i].menit*60+waktuPembalap[i].detik
		banding = waktuMulai-totalwaktu
		if mindetik > banding{
			mindetik = banding
		}
	}
	return mindetik
}