#!/usr/bin/perl
use strict;
use warnings;

open(my $rh, "<", 'data.txt') || die "Can't open file: \n";
while(read($rh, my $buf, 21)){
  my($kbn,$other)=unpack("a2a19",$buf);
  #printf "kbn=%s\n",$kbn;
  if ($kbn=="02") { $other=uc($other); }
  $buf=pack("a2a19",$kbn,$other);
  print $buf;
}
close($rh);
