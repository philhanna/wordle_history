from unittest import TestCase
from wordle import load_word_list


class Test(TestCase):

    def test_load_word_list_good(self):
        wl = load_word_list()
        self.assertIn('ATONE', wl)

    def test_load_word_list_bad(self):
        wl = load_word_list()
        self.assertNotIn('asdfi', wl)
