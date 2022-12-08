input = File.open("input.txt").read;
print input + "\n"
for i in 0..input.length
    if input[i,4].chars.to_a.uniq.length == input[i,4].length
        print "Start of packet: %d\n" % (i + 4)
        break
    end
end

for i in 0..input.length
    if input[i,14].chars.to_a.uniq.length == input[i,14].length
        print "Start of message: %d\n" % (i + 14)
        break
    end
end