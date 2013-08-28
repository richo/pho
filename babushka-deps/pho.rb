dep 'mercurial.managed' do
  provides ["hg"]
end

dep 'golang-go.managed' do
  provides ["go"]
end

dep 'libffi-dev.managed' do
  provides []
end

dep 'pho dev' do
  requires 'mercurial.managed',
           'golang-go.managed',
           'libffi-dev.managed'
end
