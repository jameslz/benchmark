function aligner_parser(args){

    var buf    = new Bytes();
    var file   = args[0] === '-' ? new File() : new File(args[0]);
    var init = '';
    
    var max_evalue    = parseFloat(args[1]);
    var min_bit_score = parseFloat(args[2]);


    while(file.readline(buf) > 0){
        
        var line = buf.toString();
        
        if (line.charAt(0) !== '#'){
            
            var t = line.split("\t");
            var query     = t[0];
            var evalue    = parseFloat(t[10]);
            var bit_score = parseFloat(t[11]);

            if( init !== query ){
                if( evalue <= max_evalue &&  bit_score >= min_bit_score){
                    print(line);
                }
                init = query;
            }
        }

    }
}

function main(args){

    if(args.length === 0){
        print("Usage:k8 best_hsp.js <alignent> <evalue> <bit_score>\n");
        exit(0);
    }
    aligner_parser(args);

}

main(arguments)