import io
import json
import pkgutil


def load_word_list():
    bs = pkgutil.get_data('wordle.data', 'word_list.json')
    with io.BytesIO(bs) as fp:
        wl = json.load(fp)
        ws = set(wl)
    return ws
