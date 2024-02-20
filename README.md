<h3>Initial settings</h3>
<ol>
<li>Download go archive: <code><b><i>wget https://dl.google.com/go/go1.22.0.linux-amd64.tar.gz</i></b></code>  </li>
<li>Extract binaries: <code><b><i>sudo tar -C /usr/local/ -xzf go1.22.0.linux-amd64.tar.gz</i></b></code></li>
<li>Set go binaries to $PATH(environment variables):

<code><b><i>
export PATH=$PATH:/usr/local/go/bin
</i></b></code>

<code><b><i>
sudo nano $HOME/.profile
</i></b></code>

<code><b><i>
source .profile
</i></b></code>
</li>

<li>Install Taskfile to run tasks: <code><b><i>go install github.com/go-task/task/v3/cmd/task@latest</i></b></code></li>
</ol>

<h3>Migrations settings</h3>
<ol>
    <li>Install goose if not installed: <code><b><i>go install github.com/pressly/goose/v3/cmd/goose@latest</i></b></code></li>
    <li>Check status of migrations: <code><b><i>task db/migrations/status</i></b></code></li>
    <li>Check version of migrations: <code><b><i>task db/migrations/version</i></b></code></li>
    <li>Validate migrations: <code><b><i>task db/migrations/validate</i></b></code></li>
    <li>Run migrations up: <code><b><i>task db/migrate/up</i></b></code></li>
</ol>

<h3>Swagger settings</h3>
<ol>
    <li>Install swagger if not installed: <code><b><i>go install github.com/swaggo/swag/cmd/swag@latest</i></b></code></li>
    <li>Generate swagger docs: <code><b><i>swag init -dir ./cmd/api/ -o ./api/docs</i></b></code></li>
</ol>