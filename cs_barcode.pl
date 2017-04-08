use strict;
use POSIX qw(ceil);

#&csvparse( sub{ <DATA> } , sub{ print map("[$_]",@{$_[0]}),"\n"; } );
#&csvparse( sub{ <DATA> } , \&printcsv );

# CONTROL CODE
use constant {
		CC1 => "a", CC2 => "b", CC3 => "c", CC4 => "d",
		CC5 => "e", CC6 => "f", CC7 => "g", CC8 => "h",
		ST => "(", SP => ")",
};

use constant MAX_LENGTH => 10;

# ALPHABET CODE
my %A_CODE = (
		A => CC1."0", B => CC1."1", C => CC1."2", D => CC1."3", E => CC1."4",
		F => CC1."5", G => CC1."6", H => CC1."7", I => CC1."8", J => CC1."9",
		K => CC2."0", L => CC2."1", M => CC2."2", N => CC2."3", O => CC2."4",
		P => CC2."5", Q => CC2."6", R => CC2."7", S => CC2."8", T => CC2."9",
		U => CC3."0", V => CC3."1", W => CC3."2", X => CC3."3", Y => CC3."4",
		Z => CC3."5",
);

# CHARACTER CODE FOR CHECK DIGIT
my %C_CODE = (
		"0" => 0, "1" => 1, "2" => 2, "3" => 3, "4" => 4,
		"5" => 5, "6" => 6, "7" => 7, "8" => 8, "9" => 9,
		"-" => 10,
		CC1() => 11, CC2() => 12, CC3() => 13, CC4() => 14,
		CC5() => 15, CC6() => 16, CC7() => 17, CC8() => 18,
);

my %CD_CODE = reverse %C_CODE;
my $DIV = keys %C_CODE;

sub csvparse{
    my ($read,$callback)=@_;
    while( defined(my $line=$read->()) ){
        for(;;){
            $line =~ s/"([^"]+)"/"\a".unpack('h*',$1)."\a"/ge;
            last unless $line =~ /"/ && defined(my $next = $read->());
            $line .= $next;
        }
        chomp $line;
        my @csv = split(/,/,$line);
        s/\a([^\a]+)\a/pack('h*',$1)/ge foreach( @csv );
        $callback->( \@csv );
    }
}

sub printcsv{
	#print map("[$_]",@{$_[0]}),"\n";
	my @a = @{$_[0]};
	print $#a+1,"===>";
	#for my $v (@{$_[0]}) {
	for my $v (@a) {
		print "[$v]";
	}
	print "\n";
}

sub justify{
	my $in=$_[0];
	if (length($in) >= MAX_LENGTH){
		return substr($in,0,MAX_LENGTH);
	}
	for (my $i=length($in) ; $i<MAX_LENGTH ; $i++){ $in .= CC4; }
	return $in
}

sub convert_alphabet{
	my $ret;
	for my $c (split //,$_[0]){
		if ($c =~ /[A-Z]/){
			$ret .= $A_CODE{$c};
		}else{
			$ret .= $c;
		}
	}
	return $ret;
}

sub parity{
	my $sum;
	for my $c (split //,$_[0]) {$sum += $C_CODE{$c};}
	return $CD_CODE{ceil($sum/$DIV)*$DIV - $sum};
}

sub cs_barcode{
	my $in=${$_[0]}[0];
	print "$in";
	if ($in !~ /^[0-9A-Z-]+$/) {
		print " => ERROR\n";
		return;
	}
	my $c=convert_alphabet($in);
	print " => $c "; 
	my $j=justify($c);
	print " => $j"; 
	my $p=parity($j);
	print " => $p"; 
	my $out=ST.$j.$p.SP;
	print " ==> $out"; 
	print "\n"; 
}

#&csvparse( sub{ <DATA> } , \&printcsv );
&csvparse( sub{ <DATA> } , \&cs_barcode );

__END__
nihon,go ahaha,ihihi
ahaha,"ihihi , ufufu","ohoho
mumumu
ahaha
gahaha"
XXXXX,"YYYYY","DDDDD,XXXXXX"
abc
+
*
&
abcAAAaaa
AAAxAAA
1
2
3
A
B
H
I
J
1-2-3ABA
1234567890-ABCDEFGHIJKLMNOPQR
0000000000
0000000001
0000000002
0000000003
0000000009
00000000-
00000000A
00000000B
00000000C
00000000I
00000000J
