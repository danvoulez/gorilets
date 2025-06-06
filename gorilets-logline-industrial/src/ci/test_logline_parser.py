from lark import Lark
import glob

def main():
    grammar = 'start: WORD\nWORD: /.+/'
    parser = Lark(grammar)
    for path in glob.glob('../logline/**/*.logline', recursive=True):
        with open(path) as f:
            parser.parse(f.read())
if __name__ == '__main__':
    main()
