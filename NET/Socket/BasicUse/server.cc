#include <sys/types.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <iostream>
#include <string.h>

using std::cout;
using std::endl;
int main() {
    int listenfd = socket(AF_INET, SOCK_STREAM, 0);
    if (listenfd == -1) {
        cout << "create socket failed" << endl;
        return -1;
    }

    struct sockaddr_in bindaddr;
    bindaddr.sin_family = AF_INET;      //协议
    bindaddr.sin_addr.s_addr = htonl(INADDR_ANY); //用户IP
    bindaddr.sin_port = htons(3000); //监听端口
    if (bind(listenfd, (struct sockaddr*)&bindaddr, sizeof bindaddr) < 0) {
        cout << "bind listen socket error" << endl;
        return -1;
    }
    if (listen(listenfd, SOMAXCONN) < 0) {
        cout << "listen error" << endl;
        return -1;
    }

    while (true) {
        struct sockaddr_in clientaddr;
        socklen_t clientaddrlen = sizeof clientaddr;
        int clientfd = accept(listenfd, (struct sockaddr*)&clientaddr, &clientaddrlen);
        if (clientfd != -1) {
            char recvBuf[32] = {0};
            int ret = recv(clientfd, recvBuf, 32, 0);
            if (ret > 0) {
                cout << "recv data from client, data : " << recvBuf << endl;
                ret = send(clientfd, recvBuf, strlen(recvBuf), 0);
                if (ret != strlen(recvBuf))
                    cout << "send data error" << endl;
                else
                    cout << "send data succesfuly data: " << recvBuf << endl;
            }else{
                cout << "recv data failed" << endl;
            }
            close(clientfd);
        }
    }
    close(listenfd);
    return 0;
}