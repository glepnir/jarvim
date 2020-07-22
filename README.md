<div align="center">
<img src="https://user-images.githubusercontent.com/41671631/88163611-81c1e880-cc45-11ea-80a0-75fb9ba40b87.png">
<p>Generate a module vim configruation like <strong>VIM PRO!</strong></p>
</div>

> I have been maintaining Thinkvim for a long time, and slowly it deviated from my original intention. At first I wanted
> to use it as a vim configuration template for Vimer to use, but slowly it has to accept everyone’s preferences, I have
> to get Add more language support and other plug-in configurations, thinkvim becomes more and more bloated. This is wrong.
> I think that vim should be lightweight, so I wrote this cli tool jarvim, which will generate a greateful configuration
> for you, there are some useful hacks in the generated configuration, I hope it can help you, if you are new to vim, you
> no longer need to refer to other people's configurations, you can use the generated configuration as your starting point
> This will greatly save your time.

## Install

You can download build binary file from release page https://github.com/glepnir/jarvim/releases

**MacOs brew**

```console
brew tap glepnir/jarvim
brew install jarvim
```

**Linux**

```console
curl -fLo  install.sh https://raw.githubusercontent.com/glepnir/jarvim/master/install.sh
sh install.sh
./jarvim -g
```

**Install From Source**

```go
go get github.com/glepnir/jarvim
```

## Usage

Here is a [gif](https://github.com/glepnir/jarvim/wiki) for how to use jarvim.

```
 -v to print jarvim version.
 -g to generate vim configuration.
```

## FAQ

- Why the symbols look weird in my vim ?

Make sure you have installed nerdfont font from https://www.nerdfonts.com/, Different fonts may be inconsistent in the performance of symbols.
The solution, If you use Mac with iterm2, you can set a different font for the symbol.

<center>
  <img src="https://user-images.githubusercontent.com/41671631/88161810-0c551880-cc43-11ea-9699-17150cd7813a.png" height="300", weight="300"/>
</center>

Another way I recommend you to use [kitty terminal](https://github.com/kovidgoyal/kitty), it has built-in symbol font support.Kitty support
Mac and Linux.

Normal graphics should be like this

<center>
  <img src="https://user-images.githubusercontent.com/41671631/88163626-8c7c7d80-cc45-11ea-9912-97d63e21c026.png" height="500", weight="800"/>
</center>

## Donate

Do you like jarvim? buy me a coffe 😘!

[![Support via PayPal](https://cdn.rawgit.com/twolfson/paypal-github-button/1.0.0/dist/button.svg)](https://www.paypal.me/bobbyhub)

| Wechat                                                                                                          | AliPay                                                                                                       |
| --------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------ |
| ![wechat](https://user-images.githubusercontent.com/41671631/84404718-c8312a00-ac39-11ea-90d7-ee679fbb3705.png) | ![ali](https://user-images.githubusercontent.com/41671631/84403276-1a714b80-ac38-11ea-8607-8492df84e516.png) |

## LICENSE

- MIT
