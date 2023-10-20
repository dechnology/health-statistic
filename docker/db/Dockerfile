FROM postgres:14.9

RUN localedef -i zh_TW -c -f UTF-8 -A /usr/share/locale/locale.alias zh_TW.UTF-8
ENV LANG zh_TW.utf8

# Set the timezone.
ENV TZ=Asia/Taipei\q

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone