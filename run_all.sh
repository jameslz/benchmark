#!/bin/sh
echo -e "\n"perl
time perl best_hsp.pl example_blast6out.tsv  1e-6 60 > res_perl.tsv

echo -e "\n"python
time python best_hsp.py example_blast6out.tsv  1e-6 60 > res_python.tsv

echo -e "\n"js
time k8 best_hsp.js example_blast6out.tsv  1e-6 60 >js_res.tsv
echo

echo -e "\n"c
time ./best_hsp_c example_blast6out.tsv  1e-6 60 > res_c.tsv

echo -e "\n"golang
time ./best_hsp_go example_blast6out.tsv  1e-6 60 > res_go.tsv

