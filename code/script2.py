# script2.py

def main():
    f = open("coba.txt", "w+")

    for i in range(10) :
        f.write("This is example lho %d\r\n" % (i + 1))

    f.close()


if __name__ == "__main__":
    main()
