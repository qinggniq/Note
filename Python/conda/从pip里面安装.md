



# 从PIP中安装

Have you tried:

```
pip install <package>
```

or

```
conda install -c conda-forge <package>
```

## 把PIP装的包弄到conda环境

1. Run `conda create -n venv_name` and `source activate venv_name`, where `venv_name` is the name of your virtual environment.
2. Run `conda install pip`. This will install pip to your venv directory.
3. Find your anaconda directory, and find the actual venv folder. It should be somewhere like `/anaconda/envs/venv_name/`.
4. Install new packages by doing `/anaconda/envs/venv_name/bin/pip install package_name`.

