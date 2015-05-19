# benchmark
Bioinformatics scripts

 


k8 version 0.2.1: http://sourceforge.net/projects/lh3/files/k8/k8-0.2.1.tar.bz2/download<br>

1. blast best hsp parser<br>
time python best_hsp.py  example_blast6out_2000.tsv 1e-5 60 >python.res.tsv<br>
time k8 best_hsp.js  example_blast6out_2000.tsv 1e-5 60 >js.res.tsv<br>
time python best_hsp.py  example_blast6out_2000.tsv 1e-5 60 >python.res.tsv<br>

