#!/usr/bin/env perl
use strict; use warnings;
my $path="~/touchNGo/";
my $raw_string=`xrandr | grep eDP1`;

my @data_block=split(' ',$raw_string);
my $block_length=scalar(@data_block);

if ( $block_length == 15 ) {
	system($path."rot.sh inverted");
} elsif ( $block_length == 16 ) {
	system($path."rot.sh normal");
}

system($path."touchStart");
