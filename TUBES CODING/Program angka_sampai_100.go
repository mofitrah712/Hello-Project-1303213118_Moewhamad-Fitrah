Program angka_sampai_100
Kamus
angka: integer
algoritma
    Input angka
    if angka > 100 then
        output "Angka lebih besar dari 100"
    else if angka < 100 then
        for i = angka to 100 do
            output i
        end for
    else if angka == 100 then
        output angka
    end if
end program