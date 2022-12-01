use List::Util qw/sum/;
use List::Util qw/max/;
open my $fh, '<', 'calories-input.txt' or die "Can't open file $!";
my $groups = do { local $/; <$fh> };

my @c;
my $max = 0;
while($groups =~ /(?<g>(\d+\n)+)/g){
    my $s = sum(split(/\n/, $+{g}));
    push(@c, $s);
    if($s > $max){
        $max = $s
    }
}

my @sorted = (sort {$b <=> $a} @c );

my $top = sum(@sorted[0..2]);

print "Max : $max\n";
print "Top : $top\n";