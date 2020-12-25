FROM gocv/opencv:4.5.0

RUN apt-get -qq update \
  && apt-get install -y \
    libleptonica-dev \
    libtesseract-dev \
    tesseract-ocr

# Load languages
RUN apt-get install -y -qq \
  tesseract-ocr-eng

ENV APPDIR $GOPATH/src/captcha
ADD . $APPDIR
WORKDIR $APPDIR

RUN go mod vendor
RUN go mod download

RUN go build -o main .
CMD ["./main"]