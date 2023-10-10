#! /usr/bin/python

from wordle import load_word_list

def contains_duplicates(word: str) -> bool:
    x = {}
    for letter in word:
        x[letter] = True
    return len(x) < len(word)

# From the list of all 5-letter words, get the frequency of each letter
wl = load_word_list()
freq = {}
for word in wl:
    for letter in word:
        if letter not in freq:
            freq[letter] = 0
        freq[letter] += 1

# Now read the list again and assign a weight to each word as the sum of the
# frequency counts for each letter.

ww = []   
for word in wl:
    weight = 0
    for letter in word:
        weight += freq[letter]
    t = (weight, word)
    ww.append(t)     

ww.sort(reverse=True)

with open("words.txt", "wt") as out:
    for weight, word in ww:
        if contains_duplicates(word):
            continue
        print(word, file=out)
