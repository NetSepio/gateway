FROM gitpod/workspace-full

RUN sudo apt update
RUN sudo apt-get install -y libappindicator1 fonts-liberation
RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
RUN sudo dpkg -i google-chrome*.deb; return 0;
RUN sudo apt-get install -y -f
RUN sudo dpkg -i google-chrome*.deb
RUN sudo curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sudo sh -s -- -b $(go env GOPATH)/bin