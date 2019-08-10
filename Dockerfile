FROM golang:1.8

LABEL maintainer="Hiromu Ochiai <otiai10@gmail.com>"

RUN apt-get update -qq

# You need librariy files and headers of tesseract and leptonica.
# When you miss these or LD_LIBRARY_PATH is not set to them,
# you would face an error: "tesseract/baseapi.h: No such file or directory"
RUN apt-get install -y -qq libtesseract-dev libleptonica-dev

# In case you face TESSDATA_PREFIX error, you minght need to set env vars
# to specify the directory where "tessdata" is located.
ENV TESSDATA_PREFIX=/usr/share/tesseract-ocr

# Load languages.
# These {lang}.traineddata would b located under ${TESSDATA_PREFIX}/tessdata.
RUN apt-get install -y -qq \
  tesseract-ocr-eng \
  tesseract-ocr-tha
# See https://github.com/tesseract-ocr/tessdata for the list of available languages.
# If you want to download these traineddata via `wget`, don't forget to locate
# downloaded traineddata under ${TESSDATA_PREFIX}/tessdata.

WORKDIR /go/src/app
COPY ./app .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]