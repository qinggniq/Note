#include <sys/types.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <iostream>
#include <string.h>

using std::cout;
using std::endl;

const char SERVER_ADDRESS[] = "127.0.0.1";
const char SEND_DATA[] = "hello word";
const  int SERVER_PORT = 3000;

int main() {
    int clientfd = socket(AF_INET, SOCK_STREAM, 0);
    if (clientfd == -1) {
        cout << "create socket failed" << endl;
        return -1;
    }

    struct sockaddr_in serveraddr;
    serveraddr.sin_family = AF_INET;
    serveraddr.sin_addr.s_addr = inet_addr(SERVER_ADDRESS);
    serveraddr.sin_port = htons(SERVER_PORT);

    if (connect(clientfd, (struct sockaddr*)&serveraddr, sizeof serveraddr) < 0) {
        cout << "connect socket error" << endl;
        return -1;
    }

    int ret = send(clientfd, SEND_DATA, strlen(SEND_DATA), 0);
    if (ret != strlen(SEND_DATA)) {
        cout << "send data faile" << endl;
        return -1;
    }
    cout << "send data successfully, data:" << SEND_DATA << endl;

    char recvBuf[32] = {0};
    ret = recv(clientfd, recvBuf, 32, 0);
    if (ret > 0) {
        cout << "successfully" << endl;
    }else{
        cout << "failed" << endl;
    }

    close(clientfd);
    return 0;

}