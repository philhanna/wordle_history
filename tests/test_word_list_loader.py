from unittest import TestCase
from wordle import load_word_list


class Test(TestCase):

    def test_load_word_list_good(self):
        ws = load_word_list()
        self.assertIn('ATONE', ws)

    def test_load_word_list_bad(self):
        ws = load_word_list()
        self.assertNotIn('asdfi', ws)
