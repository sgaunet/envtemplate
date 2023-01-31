FROM scratch AS final
LABEL maintainer="Sylvain Gaunet <sgaunet@gmail.com>"
WORKDIR /usr/bin
COPY envtemplate .
