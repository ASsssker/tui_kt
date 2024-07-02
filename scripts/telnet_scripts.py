import telnetlib as tb
import sys

_, username, password, *_ = [i.encode("ascii") for i in sys.argv]
host = "localhost"
path = b"~/code/t_kt"
end = b"\r\n"


def auth(conn: tb.Telnet, username: bytes, password: bytes) -> None:
    conn.read_until(b": ")
    conn.write(username + end)
    conn.read_until(b": ")
    conn.write(password + end)
    conn.read_until(b"$")


def readDir(path: bytes, host: str, port=23) -> list[str]:
    with tb.Telnet(host, port, timeout=4) as conn:
        auth(conn, username, password)
        conn.write(b"cd " + path + end)
        conn.read_until(b"$").decode("utf-8")
        conn.write(b"ls -l" + end)
        data = conn.read_until(b"$").decode("utf-8")
    
    data = [line for line in data.split("\r\n")]
    out = data[2: len(data)-1]

    return out


print("\n".join(readDir(path, host)))