#!/usr/bin/perl -w

use strict;
use warnings;

die "Usage:perl $0  <alignment> <evalue>  <bit_score>" if(@ARGV != 3);

my ($alignment, $max_evalue, $min_bit_score) = @ARGV;

load_alignment();

exit;

sub load_alignment{

    open(PARSER,      $alignment) || die "$!";
    
    my $init       = "";

    while(<PARSER>){
            
            chomp;
            next if(/^#/);
            my @its = split /\t/, $_;
            
            next if( $init eq $its[0] );

            my ($query, $e_value, $bit_score) = ($its[0], $its[10], $its[11]);
             
            $init  = $query;

            next if( $e_value > $max_evalue || $bit_score < $min_bit_score );
            print $_, "\n";

    }
    close PARSER;

}