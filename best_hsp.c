#include <zlib.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#include "kseq.h"
#include "kstring.h"
#include "kvec.h"

KSTREAM_INIT(gzFile, gzread, 16384)

int main(int argc, char const *argv[]){
    
    gzFile     fp; 
    kstream_t  *ks;
    
    kstring_t  str = {0, 0, 0};
    char   buf[1024];
    int    flag, min_bit_score;
    double max_e_value; 

    if(argc != 4){
        printf("%s\n", "Usage:best_hsp <tsv> <e_value> <bit_score>\n");
        return 1;
    }

    fp            = gzopen(argv[1], "r");
    max_e_value   = atof(argv[2]);    
    min_bit_score = atoi(argv[3]);

    ks            = ks_init(fp);

         
    while( ks_getuntil( ks, '\n', &str, 0) >=  0){

        if(str.l == 0 || str.s[0] == '#') continue;
        
        ks_tokaux_t aux;
        double e_value; /* e_value: 10. */
        int k, bit_score; /* bit_score: 11. */
        
        char *p , *q;
        q = strdup(str.s);

        for (p = kstrtok(str.s, "\t", &aux), k = 0; p; p = kstrtok(0, 0, &aux), ++k){
                *(char*)aux.p = 0;

                if(k == 0){
                    if( strcmp(buf, p) == 0){ 
                        flag = 0; break;
                    }else{
                        sprintf(buf,"%s", p);
                        buf[strlen(p)] = '\0';
                        flag  = 1;
                    }

                }else if(k == 10){
                    e_value   = atof(p);
                }else if(k == 11){
                    bit_score = atoi(p);                
                }
        }

        if(e_value  <= max_e_value && bit_score >= min_bit_score && flag == 1 ){
            printf("%s\n", q);
        }
    }

    ks_destroy(ks);
    gzclose(fp);
    free(str.s);
    return 0;
}