import System.IO  
import Prelude
import Data.Set
import Data.Char(ord,isUpper)
import Text.Printf

main = do  
        contents <- readFile "input.txt"
        let elves = bags contents
        let badgeSum = sum(values [commonBadge x | x <- groupElves 3 elves])
        let score = sum(values [uncurry common x | x <- compartments (bags contents)])

        printf "Badge Total %d\n" badgeSum
        printf "Common Total %d\n" score

groupElves :: Int -> [a] -> [[a]]
groupElves n = takeWhile (not.Prelude.null) . Prelude.map (Prelude.take n) . iterate (Prelude.drop n)

values :: [Char] -> [Int]
values chars = [value x | x <- chars ]

value :: Char -> Int
value v = if isUpper v then ord v - ord 'A' + 27 else ord v - ord 'a' + 1

common :: String -> String -> Char
common x y = head(toList(fromList(x) `intersection` fromList(y)))

commonBadge :: [String] -> Char
commonBadge = findChar

findChar :: [String] -> Char
findChar x = head(toList(foldedIntersect(x)))

foldedIntersect :: [String] -> Set Char
foldedIntersect x = foldl1 intersection [fromList(y) | y <- x]

bags :: String -> [String]
bags = lines

compartments :: [String] -> [(String,String)]
compartments c = [Main.split x | x <- c]

split :: [a] -> ([a], [a])
split myList = Prelude.splitAt (((length myList) + 1) `div` 2) myList