from wordle import load_word_list

def test_load_word_list_good():
    wl = load_word_list()
    assert 'ATONE' in wl

def test_load_word_list_bad():
    wl = load_word_list()
    assert 'asdfi' not in wl
